This empty directory is needed so that letsencrypt
can renew an SSL certificates. To get a certificate,
run from the parent directory (not from here):

    cd ..
    sudo certbot certonly --webroot -w $PWD -d dosa.authbank.com

To renew:

    sudo certbot renew

The certificates live in:

    /etc/letsencrypt/live/dosa.authbank.com
