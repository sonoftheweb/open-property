FROM php:8.3-fpm-alpine

# Add the www.conf configuration for PHP-FPM
ADD ./docker/api/www.conf /usr/local/etc/php-fpm.d/www.conf

# Create a new group and user for Laravel
RUN addgroup -g 1000 laravel && adduser -G laravel -g laravel -s /bin/sh -D laravel

# Create the /var/www directory
RUN mkdir -p /var/www

# Add the application code to the /var/www directory
ADD ./api /var/www

# Install PostgreSQL dev libraries and then install the PHP extensions
RUN apk add --no-cache postgresql-dev \
    && docker-php-ext-install pdo pdo_pgsql

# Change the ownership of the /var/www directory to the 'laravel' user and group
RUN chown -R laravel:laravel /var/www \
    && chown -R laravel:laravel /var/www/storage \
    && chown -R laravel:laravel /var/www/bootstrap/cache

# Set correct permissions for writing to these directories
RUN chmod -R 775 /var/www/storage \
    && chmod -R 775 /var/www/bootstrap/cache