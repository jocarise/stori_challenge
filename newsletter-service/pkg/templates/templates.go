package templates

const FooterTemplate = `<div style="font-family: Arial, sans-serif; font-size: 12px; color: #333; text-align: center; padding: 10px; border-top: 1px solid #ddd; margin-top: 20px;">
<p style="margin: 5px 0;">&copy; 2024 Stori Newsletter </p>
<p style="margin: 5px 0; color: #888;">
	<a href="%s" style="color: #007BFF; text-decoration: none;">Unsuscribe only from all newsletters</a>
	
</p>
<p style="margin: 5px 0; color: #888;">
	<a href="%s" style="color: #007BFF; text-decoration: none;">Unsuscribe from this newsletter category</a>
</p>
</div>`

const DefaultTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Email Template</title>
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; margin: 0; padding: 20px; background-color: #f4f4f4;">
    <div style="max-width: 600px; margin: auto; background: #fff; padding: 20px; border-radius: 5px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);">
        <h1 style="margin-top: 0;">Hello,</h1>
        <p>We hope this message finds you well. This is a generic email to keep you updated on our latest news and offerings.</p>
        <p>If you have any questions or need further information, feel free to reach out!</p>
        <p>Thank you for your continued support!</p>
        <p>Best regards,<br>Stori</p>
    </div>
</body>
</html>`
