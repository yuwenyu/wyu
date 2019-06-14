module github.com/yuwenyu/kernel

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.39.0

	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.4

	go.etcd.io/etcd => github.com/etcd-io/etcd v3.3.13+incompatible

	go.opencensus.io => github.com/census-instrumentation/opencensus-go v0.22.0
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.1.0
	go.uber.org/zap => github.com/uber-go/zap v1.10.0
	golang.org/x/build => github.com/golang/build v0.0.0-20190530221331-2759dfe1c117
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/exp/errors => github.com/golang/exp/errors v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190507092727-e4e5bf290fec
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190509164839-32b2708ab171
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190501051839-6835260b7148
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190502175342-a43fa875dd82
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190503185657-3b6f9c0030f7
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190513163551-3ee3066db522

	google.golang.org/api => github.com/google/google-api-go-client v0.4.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190502173448-54afdca5d873
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1

	gopkg.in/asn1-ber.v1 => github.com/go-asn1-ber/asn1-ber v0.0.0-20181015200546-f715ec2f112d
	gopkg.in/fsnotify.v1 => github.com/Jwsonic/recinotify v0.0.0-20151201212458-7389700f1b43
	gopkg.in/gorethink/gorethink.v4 => github.com/rethinkdb/rethinkdb-go v4.0.0+incompatible
	gopkg.in/inf.v0 => github.com/go-inf/inf v0.9.1
	gopkg.in/ini.v1 => github.com/go-ini/ini v1.42.0
	gopkg.in/pipe.v2 => github.com/go-pipe/pipe v0.0.0-20140414041502-3c2ca4d52544
	gopkg.in/src-d/go-billy.v4 => github.com/src-d/go-billy v4.2.0+incompatible
	gopkg.in/src-d/go-git-fixtures.v3 => github.com/src-d/go-git-fixtures v3.5.0+incompatible
	gopkg.in/tomb.v1 => github.com/go-tomb/tomb v0.0.0-20141024135613-dd632973f1e7
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v2.1.0+incompatible

	k8s.io/api => github.com/kubernetes/api v0.0.0-20190512063542-eae0ddcf85ba
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.0.0-20190514012558-1f207b29b441
	k8s.io/client-go => github.com/kubernetes/client-go v11.0.0+incompatible
	k8s.io/gengo => github.com/kubernetes/gengo v0.0.0-20190327210449-e17681d19d3a
	k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0
	k8s.io/kube-openapi => github.com/kubernetes/kube-openapi v0.0.0-20190510232812-a01b7d5d6c22
	k8s.io/utils => github.com/kubernetes/utils v0.0.0-20190506122338-8fab8cb257d5

	labix.org/v2/mgo => github.com/go-mgo/mgo v0.0.0-20180705113738-7446a0344b78
	launchpad.net/gocheck => github.com/go-check/check v0.0.0-20180628173108-788fd7840127
	sigs.k8s.io/structured-merge-diff => github.com/kubernetes-sigs/structured-merge-diff v0.0.0-20190426204423-ea680f03cc65
	sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0
)

require (
	github.com/gin-contrib/multitemplate v0.0.0-20190528082104-30e424939505
	github.com/gin-gonic/gin v1.4.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/syyongx/ii18n v0.0.0-20190531015407-03d063505fc9
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/ini.v1 v1.0.0-00010101000000-000000000000
)
