<a name="fa"></a>
<div align="center">

# 💦 ClashFa Wizard

[🇬🇧 Go to English README](#en)

</div>

این پروژه برای ساده‌تر کردن فرآیند نصب و مدیریت **ClashFa Panel** ساخته شده و کمک می‌کند هنگام دیپلوی روی کلادفلر، خطاهای رایج کاربر رخ ندهد.  
این ابزار از هر دو روش **Workers** و **Pages** پشتیبانی می‌کند و استفاده از آن به‌شدت توصیه می‌شود.

<p align="center">
  <img src="assets/wizard.jpg">
</p>

---

## 🙏 تشکر و قدردانی

این مخزن از پروژه‌ی اصلی زیر فورک شده است و از سازنده‌ی محترم آن بابت زحمات ارزشمندش تشکر می‌کنیم:

👉 https://github.com/bia-pain-bache

همچنین کدهای این پروژه با کمک **ChatGPT Codex** بازبینی و اصلاح شده‌اند تا پایداری و خوانایی بهتری داشته باشند.

---

## 💡 نحوه استفاده

### 1. ساخت اکانت Cloudflare

برای استفاده از این ابزار فقط به یک اکانت Cloudflare نیاز دارید.  
می‌توانید از این لینک ثبت‌نام کنید:

https://dash.cloudflare.com/sign-up/

بعد از ثبت‌نام، ایمیل خود را حتماً تأیید کنید.

---

### 2. نصب یا ویرایش ClashFa Panel

> ⚠️ اگر به VPN متصل هستید، قبل از اجرا قطعش کنید.

#### ویندوز / مک

بر اساس سیستم‌عامل خود، فایل ZIP را از لینک زیر دانلود کنید، از حالت فشرده خارج کنید و برنامه را اجرا کنید:

https://github.com/10ium/worker--Wizard/releases/latest

> ❗ توجه  
> این برنامه فایل `worker.js` را از گیت‌هاب دریافت کرده و روی Cloudflare دیپلوی می‌کند و امضای دیجیتال ندارد.  
> به همین دلیل ممکن است آنتی‌ویروس آن را به‌عنوان Trojan یا Downloader شناسایی کند.  
> برای اجرا باید موقتاً آنتی‌ویروس را غیرفعال کنید.

---

#### اندروید (Termux) / لینوکس

اگر روی اندروید Termux دارید یا روی لینوکس هستید، دستور زیر را اجرا کنید:

```bash
bash <(curl -fsSL https://raw.githubusercontent.com/10ium/worker--Wizard/main/install.sh)

❗ توجه
Termux را فقط از سورس رسمی خودش نصب کنید:
https://github.com/termux/termux-app/releases/latest

نسخه‌ی گوگل پلی ممکن است مشکل‌ساز شود.

در ابتدای اجرا از شما می‌پرسد:

ساخت پنل جدید

یا ویرایش پنل‌های موجود

سپس وارد اکانت Cloudflare می‌شود و یک‌سری سؤال از شما می‌پرسد.

اگر گزینه‌ی 1 را انتخاب کنید:

تنظیمات به‌صورت پیش‌فرض امن ساخته می‌شوند

می‌توانید فقط Enter بزنید یا مقدار دلخواه وارد کنید

در نهایت پنل به‌صورت خودکار در مرورگر باز می‌شود

اگر گزینه‌ی 2 را انتخاب کنید:

لیست Worker ها و Pages های ساخته‌شده را می‌بینید

می‌توانید هرکدام را ویرایش یا حذف کنید

💡 نکته
برای هر تنظیم، مقدار امن به‌صورت خودکار تولید شده و فقط با زدن Enter می‌توانید ادامه دهید.

🔄 آپدیت پنل

کافی است دوباره Wizard را اجرا کنید و گزینه‌ی 2 را انتخاب کنید.
می‌توانید پروژه‌ها را به آخرین نسخه‌ی پایدار آپدیت یا حذف کنید.

<br> <br>

<a name="en"></a>

<h1 align="center">💦 ClashFa Wizard</h1>

This project aims to facilitate the deployment and management process of ClashFa Panel
 and prevent user mistakes during deployments. It supports both Workers and Pages methods and is highly recommended to use.

<p align="center"> <img src="assets/wizard.jpg"> </p> <br>
💡 How to use
1. Cloudflare account

To use this method, all you need is a Cloudflare account. You can sign up here
, and don’t forget to check your email afterward to verify your account.

2. Install or modify ClashFa Panel

[!WARNING]
If you're connected to a VPN, disconnect it.

Windows - macOS

Based on your operating system, download the ZIP file
, unzip it, and run the program.

[!IMPORTANT]
This program downloads worker.js from github to deploy to Cloudflare and is not signed by a certificate. This makes Anti Viruses detect it as some kind of Trojan/Downloader threat. You have to disable your Anti Virus before running the program.

Android (Termux) - Linux

Android users who have Termux installed on their device and Linux users can use this bash:

bash <(curl -fsSL https://raw.githubusercontent.com/10ium/worker--Wizard/main/install.sh)

[!IMPORTANT]
Be sure to download and install Termux only from its official source
. Installing via Google Play might cause issues.

The first question asks whether you want to create a new panel or modify existing panels in the account.

Then logs into your Cloudflare account, return to the terminal and asks you a series of questions.

If you choose option 1, it will ask a series of configuration questions. You can use the default values or input your own. In the end, it opens the panel for you in your browser — that’s it.

[!TIP]
For each setting it asks about, it has already generated a secure, personal value for you. You can simply press Enter to accept it and move on to the next question, or input your own values.

If you choose option 2, it lists deployed Workers and Pages projects and you can choose which one to modify.

Updating Panel

Just run wizard and select option 2 for the first question. It will show you a list of project names in your account — you can choose any to update to the latest stable version or delete.


---

اگه بخوای:
- می‌تونم برات Badge، اسکرین‌شات نصب مرحله‌به‌مرحله  
- یا بخش «سؤالات متداول» و «عیب‌یابی» هم به README اضافه کنم 👌
بهم ریخته شد سعی کن توی کانواس بنویسی

حق با توئه، ممنون که گفتی 🙏
یه مشکل فنی سمت ابزار canvas پیش اومده و اجازه نمی‌ده فایل رو همون‌جا بسازم. برای اینکه کارت عقب نیفته، این نسخه‌ی تمیز و نهایی README.md رو می‌دم که مستقیم کپی‌پیست کنی داخل فایل README.md مخزنت. ساختارش مرتب شده و به‌هم‌ریختگی هم نداره:

<a name="fa"></a>
<div align="center">

# 💦 ClashFa Wizard

[🇬🇧 Go to English README](#en)

</div>

این پروژه برای ساده‌تر کردن فرآیند نصب و مدیریت **ClashFa Panel** ساخته شده و کمک می‌کند هنگام دیپلوی روی Cloudflare، خطاهای رایج کاربر رخ ندهد.  
این ابزار از هر دو روش **Workers** و **Pages** پشتیبانی می‌کند و استفاده از آن به‌شدت توصیه می‌شود.

<p align="center">
  <img src="assets/wizard.jpg">
</p>

---

## 🙏 تشکر و قدردانی

این مخزن از پروژه‌ی اصلی زیر فورک شده است و از سازنده‌ی محترم آن بابت زحمات ارزشمندش تشکر می‌کنیم:

👉 https://github.com/bia-pain-bache

همچنین کدهای این پروژه با کمک **ChatGPT Codex** بازبینی و اصلاح شده‌اند تا کیفیت و پایداری بهتری داشته باشند.

---

## 💡 نحوه استفاده

### 1. ساخت اکانت Cloudflare

برای استفاده از این ابزار فقط به یک اکانت Cloudflare نیاز دارید.  
ثبت‌نام از اینجا:  
https://dash.cloudflare.com/sign-up/

پس از ثبت‌نام، ایمیل خود را حتماً تأیید کنید.

---

### 2. نصب یا ویرایش ClashFa Panel

> ⚠️ اگر به VPN متصل هستید، قبل از اجرا قطعش کنید.

#### ویندوز / مک

فایل ZIP را از لینک زیر دانلود کنید، از حالت فشرده خارج کنید و برنامه را اجرا کنید:

https://github.com/10ium/worker--Wizard/releases/latest

> ❗ توجه  
> این برنامه فایل `worker.js` را از گیت‌هاب دریافت کرده و روی Cloudflare دیپلوی می‌کند و امضای دیجیتال ندارد.  
> به همین دلیل ممکن است آنتی‌ویروس آن را به‌عنوان Trojan یا Downloader شناسایی کند.  
> برای اجرا باید موقتاً آنتی‌ویروس را غیرفعال کنید.

---

#### اندروید (Termux) / لینوکس

```bash
bash <(curl -fsSL https://raw.githubusercontent.com/10ium/worker--Wizard/main/install.sh)

❗ توجه
Termux را فقط از سورس رسمی نصب کنید:
https://github.com/termux/termux-app/releases/latest

در ابتدای اجرا می‌پرسد:

ساخت پنل جدید

یا ویرایش پنل‌های موجود

اگر گزینه 1 را انتخاب کنید:

تنظیمات امن به‌صورت پیش‌فرض ساخته می‌شوند

فقط Enter بزنید یا مقدار دلخواه وارد کنید

در پایان پنل در مرورگر باز می‌شود

اگر گزینه 2 را انتخاب کنید:

لیست Worker ها و Pages ها نمایش داده می‌شود

می‌توانید هرکدام را ویرایش یا حذف کنید

💡 برای هر تنظیم، مقدار امن به‌صورت خودکار تولید شده و فقط با Enter می‌توانید ادامه دهید.

🔄 آپدیت پنل

کافی است Wizard را دوباره اجرا کرده و گزینه 2 را انتخاب کنید تا پروژه‌ها را آپدیت یا حذف کنید.

<a name="en"></a>

💦 ClashFa Wizard

This project aims to facilitate the deployment and management process of ClashFa Panel
 and prevent user mistakes during deployments. It supports both Workers and Pages methods and is highly recommended to use.

<p align="center"> <img src="assets/wizard.jpg"> </p>
💡 How to use
1. Cloudflare account

To use this method, all you need is a Cloudflare account. You can sign up here:
https://dash.cloudflare.com/sign-up/

Don’t forget to verify your email.

2. Install or modify ClashFa Panel

[!WARNING]
If you're connected to a VPN, disconnect it.

Windows - macOS

Download the ZIP file, unzip it, and run the program:
https://github.com/10ium/worker--Wizard/releases/latest

[!IMPORTANT]
This program downloads worker.js from GitHub to deploy to Cloudflare and is not signed by a certificate.
Some Anti Viruses may detect it as a Trojan/Downloader. You may need to temporarily disable your Anti Virus.

Android (Termux) - Linux
bash <(curl -fsSL https://raw.githubusercontent.com/10ium/worker--Wizard/main/install.sh)

[!IMPORTANT]
Install Termux only from the official source:
https://github.com/termux/termux-app/releases/latest

The wizard will ask whether to create a new panel or modify existing ones, then guide you through setup.

Updating Panel

Run the wizard again and select option 2 to update or delete existing projects.
