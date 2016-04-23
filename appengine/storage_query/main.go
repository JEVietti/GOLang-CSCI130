package storage_query


import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"encoding/base64"
)

func init() {
	http.HandleFunc("/", handler)
}

const gcsBucket = "deal-breaker.appspot.com"


type gcsDelimit struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

func handler(res http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler NewClient: ", err)
		return
	}
	defer client.Close()

	d := &gcsDelimit{
		ctx:    ctx,
		res:    res,
		client: client,
		bucket: client.Bucket(gcsBucket),
	}

	d.createFiles()
	d.listFiles()
	io.WriteString(d.res, "\nResults from files WITH delimiter\n")
	d.listDirectory("", "/", "  ")
	io.WriteString(d.res, "\nResults from files WITHOUT delimiter\n")
	d.listDirectory("", "", "  ")

}

func (d *gcsDelimit) listDirectory(name, delim, indent string) {

	query := &storage.Query{
		Prefix:    name,
		Delimiter: delim,
	}

	for query != nil {
		objs, err := d.bucket.List(d.ctx, query)
		if err != nil {
			log.Errorf(d.ctx, "listBucketDirMode: unable to list bucket %q: %v", gcsBucket, err)
			return
		}
		query = objs.Next

		for _, obj := range objs.Results {
			fmt.Fprintf(d.res, "%v%v\n", indent, obj.Name)
		}

		fmt.Fprintf(d.res, "%v\n", objs.Prefixes)

		for _, pfix := range objs.Prefixes {
			log.Infof(d.ctx, "DIR: %v", pfix)
			d.listDir(pfix, delim, indent+"  ")
		}
	}
}

func (d *gcsDelimit) listFiles() {
	io.WriteString(d.res, "\nRetrieving file names...\n")

	client, err := storage.NewClient(d.ctx)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}
	defer client.Close()

	objs, err := client.Bucket(gcsBucket).List(d.ctx, nil)
	if err != nil {
		log.Errorf(d.ctx, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		io.WriteString(d.res, obj.Name+"\n")
	}
}

func (d *gcsDelimit) createFiles() {
	idx := 0
	io.WriteString(d.res, "\nCreating more files for listbucket...\n")
	for _, n := range []string{"test"} {
		d.createFile(n, idx)
		idx = (idx + 1) & 3
	}
}

func generateUUID(ctx context.Context, fileName string) string {
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Errorf(ctx, "createUUID: unable to write uuid data to bucket %q, file %q: %v", gcsBucket, fileName, err)
	}
	return uuid.String()
}

func (d *gcsDelimit) createFile(fileName string, qid int) {
	fmt.Fprintf(d.res, "Creating file /%v/%v\n", gcsBucket, fileName)

	wc := d.bucket.Object(fileName).NewWriter(d.ctx)
	wc.ContentType = "text/plain"
	b64 := base64.URLEncoding.EncodeToString([]byte("Hello, World!"))

	wc.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}

	if _, err := wc.Write([]byte("Header  ")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if _, err := wc.Write([]byte(generateUUID(d.ctx, fileName) + "\n\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return

	}
	if _, err := wc.Write([]byte("Query:  " + "\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if _, err := wc.Write([]byte("Hello, World!" + "\n\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}

	if _, err := wc.Write([]byte("Query encoded  " + "\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if _, err := wc.Write([]byte(b64 + "\n\n")); err != nil {
		log.Errorf(d.ctx, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
	if err := wc.Close(); err != nil {
		log.Errorf(d.ctx, "createFile: unable to close bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
}
