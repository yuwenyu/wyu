# wyu

**Skeleton for Golang up-v1.12.*** `v1.0.0`

## **_Directory Constructure_**
~~~
app
 |- http
    |- controllers
 |- api/
  - main.go
resources
 |- assets
 |- lang
 |- views
    |- layout
  - favicon.ico
routes
  - http.go
storage
 |- logs/
    |- *
     - *.log
vendor/
 |- * [Source for Golang]
 |- github.com
    |- yuwenyu
       |- kernel
          |- *
        - kernel.go
        - go.mod
- autoload.go
- go.mod
- go.sum
~~~

## **_Nginx Configure_**
~~~
upstream goserver_zeroplan {
    server 127.0.0.1:8888;
}

server {
    charset utf-8;
    client_max_body_size 128M;

    #listen 80; 
    #listen [::]:80 default_server ipv6only=on; 

    server_name zeroplango-dev.com www.zeroplango-dev.com;
  
    location / {
        try_files $uri @backend;
    }

    #location /(css|js|fonts|img)/ {
    #    access_log off;
    #    expires 1d;
    #    root "/path/to/app_b/static";
    #    try_files $uri @backend;
    #}

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;
        proxy_pass http://goserver_zeroplan;
    }

    #location ~ \.(js|css|png|jpg|gif|swf|ico|pdf|mov|fla|zip|rar)$ {
    #    try_files $uri =404;
    #}
    #error_page 404 /404.html;

    #location ~ /\.(ht|svn|git) {
    #    deny all;
    #}
}
~~~



