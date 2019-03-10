# notify-mail
Send mail notifications via google's gmail

# How does it work

It uses a gmail account to send notification mails.  

# How to Install

## Via package

Download latest version from https://github.com/adelolmo/notify-mail/releases/latest

```
sudo dpkg -i notify-mail_1.2.0_amd64.deb
```

## Via debian/ubuntu repository

### Setup repository

```
sudo apt-get install apt-transport-https
```

For amd64:
```
wget -O - http://adelolmo.github.io/andoni.delolmo@gmail.com.gpg.key | sudo apt-key add -
echo "deb http://adelolmo.github.io xenial main" | sudo tee /etc/apt/sources.list.d/adelolmo.list
sudo apt-get update
```
For arm:
```
wget -O - http://adelolmo.github.io/andoni.delolmo@gmail.com.gpg.key | sudo apt-key add -
echo "deb http://adelolmo.github.io jessie main" | sudo tee /etc/apt/sources.list.d/adelolmo.list
sudo apt-get update
```

### Install package
```
sudo apt-get install notify-mail
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
sudo apt-get purge notify-mail
```