## Starting container:

```sh
docker run -d                                   \
--mount type=bind,source=/root/data,target=/db  \
--restart unless-stopped                        \
-e SATANA_SELECTED='false'                      \
-e RESELECT_SATANA='false'                      \
-e DB_FILE='/db/app.db'                         \
-p 8080:8080                                    \
secret-satana
```


## Setingup for nginx
1. Clone proj
1. Create cersts

> sudo certbot certonly -d <domain> -d www.<domain> -i nginx

1. Create Virtual Configuration or use existing

> nano /secret-santa/<domain>.conf

Successfully received certificate.
Certificate is saved at: /etc/letsencrypt/live/secret-satana.online/fullchain.pem
Key is saved at:         /etc/letsencrypt/live/secret-satana.online/privkey.pem
This certificate expires on 2025-03-19.
These files will be updated when the certificate renews.
Certbot has set up a scheduled task to automatically renew this certificate in the background.

1. Create Symlink to /etc/nginx/sites-available

> ln -s /secret-santa/<domain>.conf /etc/nginx/sites-enabled/

1. Check 

> sudo nginx -t

1. Restart 

> systemctl restart nginx