# notify-mail
Send mail notifications via google's gmail

# How does it work

It uses a gmail account to send notification mails.  

# How to Install

## Via package

Download latest version from https://github.com/adelolmo/notify-mail/releases/latest

```
# dpkg -i notify-mail_1.2.0_amd64.deb
```

## Via debian/ubuntu repository

### Setup repository

Follow the instructions [here](https://adelolmo.github.io).

### Install package
```
# apt-get install notify-mail
```

## Configuration

You will need a gmail account first of all. Make sure that you don't have "two steps authentication" enable. 

You need to configure two environment variables:
* NOTIFY_MAIL_ACCOUNT is the mail account used for sending the mails.
* NOTIFY_MAIL_PASSWORD is the account's password.

```
export NOTIFY_MAIL_ACCOUNT=direction@gmail.com
export NOTIFY_MAIL_PASSWORD=your_password
```

# How to Use

## Send with body content 

```
notify-mail -recipient="email_address@email_provider.com" -subject="Notification" -message="Your message"
```

## Send with template

```
notify-mail -recipient="email_address@email_provider.com" -subject="Notification" -template="/path/to/template.html" -variables="{{var1}}=value of var1,{{var2}}=value of var2"
```

# How to remove
```
# apt-get remove notify-mail
```