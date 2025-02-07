# Cluster config file description

``` shell
cluster_id: cluster                                 # cluster name
architecture: amd64                                 # deploy cluster architecture, support amd64 or arm64
platform: libvirt                                   # deployment platform is libvirt、openstack、pxe
                                                    # Parameters need to be set according to different deployment platforms
osImage:
  type:                                             # Specify the type of operating system, such as nestos or generalos.
username: root                                      # Specify the username for ssh login
password: $1$yoursalt$UGhjCXAJKpWWpeN8xsF.c/        # Specify the password for ssh login
sshkey: "/root/.ssh/id_rsa.pub"                     # The storage path of the ssh-key file
master:                                             # master config
- hostname: k8s-master01
  hardwareinfo:                                     
    cpu: 4
    ram: 8192                                       
    disk: 50                                        
  ip: "192.168.132.11"                              
worker:                                             # worker config
- hostname: k8s-worker01            
  hardwareinfo:
    cpu: 4
    ram: 8192
    disk: 50
  ip: ""                                            # If the worker node IP address is not set, it will be automatically assigned by dhcp and will be empty by default.
runtime: isulad                                     # support docker、isulad、containerd and crio
kubernetes:                                         
  kubernetes-version: "v1.29.1"                   
  kubernetes-apiversion: "v1beta3"                  # support v1beta3、v1beta2、v1beta1
  apiserver-endpoint: "192.168.132.11:6443"          
  image-registry: "registry.k8s.io"                 # The image repository address used during Kubeadm initialization
  registryMirror: ""                                # The mirror site address of the image repository used when downloading the container image    
  pause-image: "pause:3.9"                         
  release-image-url: ""                         
  token: ""                                         # automatically generated by default
  adminkubeconfig: /etc/nkd/cluster/admin.config    # path of admin.conf
  certificatekey: ""                                # The key used to decrypt the certificate in the downloaded Secret when adding a new control plane node
  packageList:                                      # List of RPM package names that need to be installed in the cluster environment  
  rpmPackagePath: ""                                # Path to the RPM package files that need to be installed in the cluster environment
  network:                                          
    service-subnet: "10.96.0.0/16"                  
    pod-subnet: "10.244.0.0/16"                     
    plugin: ""                                      # network plugin
housekeeper:                                                                                          # housekeeper
  deployhousekeeper: false                                                                           
  operatorimageurl: "hub.oepkgs.net/nestos/housekeeper/{arch}/housekeeper-operator-manager:{tag}"     # housekeeper-operator image URL
  controllerimageurl: "hub.oepkgs.net/nestos/housekeeper/{arch}/housekeeper-controller-manager:{tag}" # housekeeper-controller image URL  
certasset:                                          # Configure user-defined certificate file path list, automatically generated by default
  rootcacertpath: ""                
  rootcakeypath: ""
  etcdcacertpath: ""
  etcdcakeypath: ""
  frontproxycacertpath: ""
  frontproxycakeypath: ""
  sapub: ""
  sakey: ""
```

Specify deployment platform configuration parameters for libvirt as an example:
``` shell
platform: libvirt                                   # Deployment platform is libvirt
infraPlatform
  uri: qemu:///system                                
  osPath:                                           # Specify the operating system image address for deploying cluster machines, supporting architectures x86_64 or aarch64
  cidr: 192.168.132.0/24                            # Routing address
  gateway: 192.168.132.1                            # Gateway address
```

To set the deployment platform to openstack, you need to reset the "infraplatform" field configuration parameters.
``` shell
platform: openstack                                   
infraplatform                      
	username:                                           # openstack username, requires permission to create resources                                      
	password:                                           # openstack login password, used to log in to the openstack platform
	tenant_name:                                        # openstack tenant name, the collection the user belongs to, for example: admin
	auth_url:                                           # openstack auth_url，example：http://{ip}:{port}/v3
	region:                                             # Used for resource isolation, for example: RegionOne
	internal_network:                                   
	external_network:                                  
	glance_name:                                        # qcow2 image
	availability_zone:                                  # default nova
```

Specify deployment platform configuration parameters for PXE as an example:
``` shell
platform: pxe                                        # Deployment platform is PXE
infraPlatform
  ip:                                                # IP address of the HTTP server
  httpServerPort: "9080"                             # Port number of the HTTP server
  httpRootDir: /var/www/html/                        # Root directory of the HTTP server
  tftpServerPort: "69"                               # Port number of the TFTP server
  tftpRootDir: /var/lib/tftpboot/                    # Root directory of the TFTP server
```

## Image Download Links

- NestOS image download, please visit the [official website](https://nestos.openeuler.org/), and download the NestOS For Container version.
- For OpenEuler image download, please visit the [official website](https://www.openeuler.org/).

## Password Cipher Generation Methods:
- When specifying the underlying operating system of the cluster as nestos, a cipher password needs to be used. Here's the generation method:
  ``` shell
  openssl passwd -1 -salt yoursalt
  Password: qwer1234!@#$
  $1$yoursalt$UGhjCXAJKpWWpeN8xsF.c/
  ```

- When deploying the platform as pxe, a cipher password needs to be used.
  ``` shell
  # python3  
  Python 3.7.9 (default, Mar  2 2021, 02:43:11)
  [GCC 7.3.0] on linux
  Type "help", "copyright", "credits" or "license" for more information.  
  >>> import crypt  
  >>> passwd = crypt.crypt("myPasswd")  
  >>> print (passwd)  
  $6$sH1qri2n14V1VCv/$fWnV3rPv95gWHJ3wZu6o0bBGy.SnllSw4a2HuoP45jXfI9fCrwe60AULO/0aXS7dWTSwvwdqqY4yFhwUdJcb.0
  ```