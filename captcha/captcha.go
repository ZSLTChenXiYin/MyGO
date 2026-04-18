package captcha

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"text/template"
	"time"
)

const (
	CAPTCHA_DIC = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type CaptchaGenerator struct {
	length uint
	rander *rand.Rand
}

func NewCaptchaGenerator(length uint) *CaptchaGenerator {
	return &CaptchaGenerator{
		length: length,
		rander: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (g *CaptchaGenerator) Generate() string {
	captcha := make([]byte, g.length)
	for i := 0; i < int(g.length); i++ {
		captcha[i] = CAPTCHA_DIC[g.rander.Intn(len(CAPTCHA_DIC))]
	}
	return string(captcha)
}

const (
	CAPTCHA_TEMPLATE = `<!DOCTYPE html>
<html lang="zh-CN">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>您的验证码</title>
    <style>
        body {
            font-family: 'Helvetica Neue', Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            margin: 0;
            padding: 0;
            background-color: #f5f5f5;
        }

        .container {
            max-width: 600px;
            margin: 20px auto;
            padding: 20px;
            background-color: #fff;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
        }

        .header {
            text-align: center;
            padding: 20px 0;
            border-bottom: 1px solid #eee;
        }

        .logo {
            height: auto;
        }

        .content {
            padding: 30px 20px;
        }

        .verification-code {
            background: #f8f9fa;
            padding: 15px;
            text-align: center;
            margin: 25px 0;
            font-size: 28px;
            font-weight: bold;
            letter-spacing: 5px;
            color: #1a73e8;
            border-radius: 4px;
            border: 1px dashed #dadce0;
        }

        .footer {
            text-align: center;
            padding: 20px;
            font-size: 12px;
            color: #999;
            border-top: 1px solid #eee;
        }

        .button {
            display: inline-block;
            padding: 12px 24px;
            background-color: #1a73e8;
            color: white !important;
            text-decoration: none;
            border-radius: 4px;
            font-weight: bold;
            margin-top: 20px;
        }

        .note {
            font-size: 14px;
            color: #666;
            margin-top: 30px;
        }
    </style>
</head>

<body>
    <div class="container">
        <div class="header">
            <h1 class="logo">浊水楼台</h1>
        </div>

        <div class="content">
            <h2>您好！</h2>
            <p>您正在尝试进行账户验证，请输入以下验证码完成操作：</p>

            <div class="verification-code">
                {{.MyGO.Captcha}}
            </div>

            <p>该验证码将在 <strong>5分钟</strong> 后失效，请尽快使用。</p>

            <p>如果您没有请求此验证码，请忽略此邮件或联系我们的支持团队。</p>

            <div class="note">
                <p>为保障您的账户安全，请勿向他人泄露此验证码。</p>
            </div>
        </div>

        <div class="footer">
            <p>© 2026 浊水楼台. 保留所有权利</p>
            <p>关于我们</p>
            <p>
                <a href="https://github.com/BPLDuskRain" style="color: #1a73e8; text-decoration: none;">夕雨落</a> |
                <a href="https://github.com/ZSLTChenXiYin" style="color: #1a73e8; text-decoration: none;">陈汐胤</a> |
                <a href="https://github.com/xunshi123" style="color: #1a73e8; text-decoration: none;">寻世</a>
            </p>
        </div>
    </div>
</body>

</html>`
)

type CaptchaTemplateGenerator struct {
	text string
	tmpl *template.Template
}

func NewCaptchaTemplateGenerator() *CaptchaTemplateGenerator {
	return &CaptchaTemplateGenerator{
		tmpl: template.New("captcha"),
	}
}

func (g *CaptchaTemplateGenerator) Open(path string) error {
	// 读取HTML模板
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	text, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	g.tmpl = template.Must(g.tmpl.Parse(string(text)))
	g.text = string(text)

	return nil
}

func (g *CaptchaTemplateGenerator) Use(text string) {
	g.tmpl = template.Must(g.tmpl.Parse(text))
	g.text = string(text)
}

func (g *CaptchaTemplateGenerator) Generate(captcha string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(string(g.text))

	buf.Reset()

	mygo := map[string]any{
		"MyGO": map[string]string{"Captcha": captcha},
	}

	err := g.tmpl.Execute(&buf, mygo)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
