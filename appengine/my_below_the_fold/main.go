package my_below_the_fold



import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"html/template"
	"net/http"
	"os"
	"io"
	"crypto/sha1"
	"fmt"
)

var tpl *template.Template

const gcsBucket = "deal-breaker.appspot.com"

type gcsPhotos struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}


func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
	resourceHandler("css")
	resourceHandler("img")
	resourceHandler("photos")
	http.HandleFunc("/", viewPage)
}

func viewPage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	gcs := configureCloud(res, req)
	photos := gcs.retrievePhotos()
	if len(photos) == 0 {
		photos = gcs.uploadPhotos()
	}
	tpl.Execute(res, photos)
}


func resourceHandler(resourceDirectory string) {
	fs := http.FileServer(http.Dir(resourceDirectory))
	fs = http.StripPrefix("/"+resourceDirectory, fs)
	http.Handle("/"+resourceDirectory+"/", fs)
}


func configureCloud(res http.ResponseWriter, req *http.Request) (gcs *gcsPhotos) {
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler NewClient: ", err)
		return
	}
	defer client.Close()

	gcs = &gcsPhotos{
		ctx:    ctx,
		res:    res,
		client: client,
		bucket:   client.Bucket(gcsBucket),
	}
	return
}

func (gcs *gcsPhotos) retrievePhotos() (photos []string) {
	files, err := gcs.bucket.List(gcs.ctx, nil)
	if err != nil {
		log.Errorf(gcs.ctx, "listBucketDirMode: unable to get the buckets %q: %v", gcsBucket, err)
		return
	}
	for _, name := range files.Results {
		photos = append(photos, name.Name)
	}
	return
}


func (gcs *gcsPhotos) uploadPhotos() []string {
	ext := ".jpg"
	subDir := "photos/"
	testPhotos := []string{ "0","2","3","4","5","6","7","8","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24"}

	for _, name := range testPhotos {
		ffile := subDir + name + ext
		srcFile, fName, err := gcs.read(ffile)
		if err != nil {
			log.Errorf(gcs.ctx, "Open / read file error %v", err)
			return nil
		}
		fName = fName + ext
		gcs.write(fName, srcFile)
		log.Infof(gcs.ctx, "In File: %s, Out File: %s\n", ffile, fName)
	}
	return gcs.retrievePhotos()
}

func (gcs *gcsPhotos) read(filename string) (buf []byte, outFilename string, err error) {
	src, err := os.Open(filename)
	if err != nil {
		return
	}
	defer src.Close()
	size, err := src.Stat()
	if err != nil {
		return
	}
	buf = make([]byte, size.Size())
	_, err = src.Read(buf)
	if err != nil {
		return
	}
	outFilename = getSha(buf)
	return
}

func (gcs *gcsPhotos) write(filename string, data []byte) {
	writer := gcs.bucket.Object(filename).NewWriter(gcs.ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	defer writer.Close()
	writer.Write(data)
}

func getSha(data []byte) string {
	h := sha1.New()
	io.WriteString(h, string(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}