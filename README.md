## 📦 Features

### Send emails via API

Send emails through a simple API request. The API supports the `POST /api/send` method and requires the following parameters in the request body (in JSON format):

```json
{
  "to": "example@tdl.com",
  "subject": "Email subject",
  "body": "Hello, demo is awesome!"
}
```

or with template:

```json
{
  "to": "example@tdl.com",
  "subject": "Email subject",
  "templateName": "welcome",
  "content": {
    "Content": "Hello, demo is awesome!",
    "Link": "https://github.com/anuwatthisuka"
  }
}
```

#### Response example

```json
{
  "message": "Email sent successfully"
}
```

```json
{
  "error": "Failed to send email: <error message>"
}
```

---
