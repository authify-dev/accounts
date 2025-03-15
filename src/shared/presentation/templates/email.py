from core.settings import settings


def get_data_for_email_activation_success(user_name):
    subject = f"¡Bienvenido/a a bordo, {user_name}! Activación de Cuenta Exitosa"
    message = settings.EMAIL_WELCOME_TEMPLATE.replace("{user_name}", user_name)
    return subject, message


def get_data_for_email_activate_account(user_name, activation_code):
    subject = f"Activa tu cuenta, {user_name}"

    message = settings.EMAIL_ACTIVATION_TEMPLATE.replace("{user_name}", user_name).replace("{activation_code}", activation_code)

    return subject, message


def get_data_for_email_two_factor(user_name, two_factor_code):
    subject = f"Código de Seguridad para {user_name} - Autenticación de Dos Factores"
    message = f"""
<html>
<head>
    <style type="text/css">
        @import url('https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;700&display=swap');

        body {{
            font-family: 'Open Sans', sans-serif;
            margin: 0;
            padding: 0;
            background-color: #F5F5F7;
            color: #333;
        }}

        .container {{
            max-width: 600px;
            margin: 40px auto;
            background-color: #FFF;
            border-radius: 10px;
            overflow: hidden;
            box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
        }}

        .header {{
            background: linear-gradient(135deg, #004482, #0071BC);
            color: #ffffff;
            padding: 20px 30px;
            font-size: 24px;
            text-align: center;
        }}

        .content {{
            padding: 20px 30px;
            color: #666;
        }}

        strong, em {{
            color: #004482;
        }}

        ol li {{
            margin-bottom: 12px;
        }}

        .footer {{
            background-color: #EBF0F3;
            color: #666;
            padding: 20px 30px;
            text-align: center;
            font-size: 14px;
        }}
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            Seguridad en cada paso, {user_name}.
        </div>
        <div class="content">
            <p>Para garantizar la seguridad de tu cuenta, te hemos enviado este código como parte de nuestro proceso de autenticación de dos factores.</p>

            <p>Tu código de seguridad es: <strong>{two_factor_code}</strong></p>

            <p>Por favor, ingresa este código en la página de inicio de sesión para confirmar tu identidad y completar el proceso de inicio de sesión. Recuerda, este código es válido solo por un corto periodo de tiempo.</p>

            <p>La seguridad de tu cuenta es nuestra prioridad. Si no has intentado iniciar sesión recientemente y recibiste este correo, te recomendamos cambiar tu contraseña de inmediato y contactarnos.</p>

            <p>Gracias por tomar las medidas necesarias para mantener tu cuenta segura. ¡Tu tranquilidad es nuestra tranquilidad!</p>
        </div>
        <div class="footer">
            Saludos seguros,<br>
            El equipo de {settings.TEAM_NAME}
        </div>
    </div>
</body>
</html>
"""
    return subject, message
