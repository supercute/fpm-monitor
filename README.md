# PHP-FPM Monitor

Terminal-based real-time PHP-FPM Monitor

## Installation

Download latest release and run

## PHP-FPM Configuration

Enable the status page in your PHP-FPM pool configuration:

```
    location /fpm-status {
        fastcgi_pass app-app:9090;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }
```

### Help

```
Usage of ./fpm-monitor:
    -lang string
        Language: en or ru (default "en")
    -url string
        PHP-FPM status URL (default "http://localhost/status")
```