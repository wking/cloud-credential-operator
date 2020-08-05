// Code generated for package bootstrap by go-bindata DO NOT EDIT. (@generated)
// sources:
// bindata/bootstrap/cloudcredential_v1_credentialsrequest.yaml
// bindata/bootstrap/namespace.yaml
package bootstrap

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bootstrapCloudcredential_v1_credentialsrequestYaml = []byte(`apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: credentialsrequests.cloudcredential.openshift.io
spec:
  group: cloudcredential.openshift.io
  names:
    kind: CredentialsRequest
    listKind: CredentialsRequestList
    plural: credentialsrequests
    singular: credentialsrequest
  scope: Namespaced
  subresources:
    status: {}
  version: v1
  validation:
    openAPIV3Schema:
      description: CredentialsRequest is the Schema for the credentialsrequests API
      type: object
      required:
      - spec
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          description: CredentialsRequestSpec defines the desired state of CredentialsRequest
          type: object
          required:
          - secretRef
          properties:
            providerSpec:
              description: ProviderSpec contains the cloud provider specific credentials
                specification.
              type: object
            secretRef:
              description: SecretRef points to the secret where the credentials should
                be stored once generated.
              type: object
              properties:
                apiVersion:
                  description: API version of the referent.
                  type: string
                fieldPath:
                  description: 'If referring to a piece of an object instead of an
                    entire object, this string should contain a valid JSON/Go field
                    access statement, such as desiredState.manifest.containers[2].
                    For example, if the object reference is to a container within
                    a pod, this would take on a value like: "spec.containers{name}"
                    (where "name" refers to the name of the container that triggered
                    the event) or if no container name is specified "spec.containers[2]"
                    (container with index 2 in this pod). This syntax is chosen only
                    to have some well-defined way of referencing a part of an object.
                    TODO: this design is not final and this field is subject to change
                    in the future.'
                  type: string
                kind:
                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
                resourceVersion:
                  description: 'Specific resourceVersion to which this reference is
                    made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                  type: string
                uid:
                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                  type: string
        status:
          description: CredentialsRequestStatus defines the observed state of CredentialsRequest
          type: object
          required:
          - lastSyncGeneration
          - provisioned
          properties:
            conditions:
              description: Conditions includes detailed status for the CredentialsRequest
              type: array
              items:
                description: CredentialsRequestCondition contains details for any
                  of the conditions on a CredentialsRequest object
                type: object
                required:
                - status
                - type
                properties:
                  lastProbeTime:
                    description: LastProbeTime is the last time we probed the condition
                    type: string
                    format: date-time
                  lastTransitionTime:
                    description: LastTransitionTime is the last time the condition
                      transitioned from one status to another.
                    type: string
                    format: date-time
                  message:
                    description: Message is a human-readable message indicating details
                      about the last transition
                    type: string
                  reason:
                    description: Reason is a unique, one-word, CamelCase reason for
                      the condition's last transition
                    type: string
                  status:
                    description: Status is the status of the condition
                    type: string
                  type:
                    description: Type is the specific type of the condition
                    type: string
            lastSyncGeneration:
              description: LastSyncGeneration is the generation of the credentials
                request resource that was last synced. Used to determine if the object
                has changed and requires a sync.
              type: integer
              format: int64
            lastSyncTimestamp:
              description: LastSyncTimestamp is the time that the credentials were
                last synced.
              type: string
              format: date-time
            providerStatus:
              description: ProviderStatus contains cloud provider specific status.
              type: object
            provisioned:
              description: Provisioned is true once the credentials have been initially
                provisioned.
              type: boolean
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`)

func bootstrapCloudcredential_v1_credentialsrequestYamlBytes() ([]byte, error) {
	return _bootstrapCloudcredential_v1_credentialsrequestYaml, nil
}

func bootstrapCloudcredential_v1_credentialsrequestYaml() (*asset, error) {
	bytes, err := bootstrapCloudcredential_v1_credentialsrequestYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/cloudcredential_v1_credentialsrequest.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrapNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  annotations:
    openshift.io/node-selector: ""
  labels:
    controller-tools.k8s.io: "1.0"
    openshift.io/cluster-monitoring: "true"
  name: openshift-cloud-credential-operator
`)

func bootstrapNamespaceYamlBytes() ([]byte, error) {
	return _bootstrapNamespaceYaml, nil
}

func bootstrapNamespaceYaml() (*asset, error) {
	bytes, err := bootstrapNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bootstrap/cloudcredential_v1_credentialsrequest.yaml": bootstrapCloudcredential_v1_credentialsrequestYaml,
	"bootstrap/namespace.yaml":                             bootstrapNamespaceYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"bootstrap": {nil, map[string]*bintree{
		"cloudcredential_v1_credentialsrequest.yaml": {bootstrapCloudcredential_v1_credentialsrequestYaml, map[string]*bintree{}},
		"namespace.yaml": {bootstrapNamespaceYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
