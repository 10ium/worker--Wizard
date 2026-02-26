<div dir="rtl" align="right">

# ClashFa Wizard

ابزار **ClashFa Wizard** برای دیپلوی و مدیریت Worker/Pages روی Cloudflare ساخته شده تا فرآیند راه‌اندازی ساده‌تر شود و خطاهای رایج کمتر شوند.

> این پروژه یک فورک/نسخه سفارشی‌شده است و متناسب با نیازهای ClashFa به‌روزرسانی شده است.

---

## ✨ امکانات

- دیپلوی با دو روش:
  - **Cloudflare Workers**
  - **Cloudflare Pages**
- انتخاب سورس Worker در زمان اجرا:
  - Worker اصلی (Legacy)
  - لیست Workerهای پیش‌فرض پروژه
  - لینک سفارشی کاربر
- آپلود نهایی همیشه با نام `worker.js` برای سازگاری با Cloudflare
- پشتیبانی از حالت **Legacy** فقط برای Worker اصلی:
  - اعمال تنظیمات و Environmentهای قدیمی
  - برگرداندن لینک نهایی با `/panel`
- برای Workerهای جدید/سفارشی:
  - دیپلوی ساده‌تر
  - بدون وابستگی به تنظیمات Legacy

---

## 🔗 Worker Sourceهای پیش‌فرض

- `https://raw.githubusercontent.com/10ium/free-config/main/worker/iptv_player.txt`
- `https://raw.githubusercontent.com/10ium/free-config/main/worker/ClashFa_Mirror_Pro.txt`
- `https://raw.githubusercontent.com/10ium/free-config/refs/heads/main/worker/great_mihomo_converter`
- `https://raw.githubusercontent.com/10ium/free-config/main/worker/iran_proxy.txt`

> در حالت **Pages**، لینک‌های پیش‌فرض فقط برای سناریوی تعریف‌شده نمایش داده می‌شوند و امکان وارد کردن لینک سفارشی هم وجود دارد.

---

## 🧩 پیش‌نیازها

- یک اکانت Cloudflare
- اتصال اینترنت پایدار
- قطع بودن VPN (در صورت بروز مشکل DNS/Login)

---

## 🚀 نصب و اجرا

### Windows / macOS

آخرین نسخه را از بخش Release همین مخزن دانلود و اجرا کنید.

### Linux / Android (Termux)

```bash
bash <(curl -fsSL https://raw.githubusercontent.com/10ium/worker--Wizard/main/install.sh)
```

---

## 🛠️ نحوه استفاده

1. برنامه را اجرا کنید.
2. بین **Create** یا **Modify** انتخاب کنید.
3. روش دیپلوی (**Workers** یا **Pages**) را انتخاب کنید.
4. سورس Worker را انتخاب کنید (پیش‌فرض/Legacy/Custom).
5. در صورت انتخاب Legacy، تنظیمات اضافی همان حالت پرسیده می‌شود.
6. دیپلوی انجام می‌شود و لینک نهایی تحویل داده می‌شود.

---

## 🧪 Build برای توسعه‌دهنده

```bash
go build ./...
go test ./...
```

ساخت باینری خروجی (مثال):

```bash
make build VERSION=$(cat VERSION) GOOS=linux GOARCH=amd64
```

---

## 🙏 تشکر

از سازنده اصلی پروژه BPB Wizard بابت ایده و پیاده‌سازی اولیه بسیار متشکریم.

این فورک برای سناریوهای ClashFa بازطراحی/اصلاح شده و بخشی از بهینه‌سازی‌ها با کمک **ChatGPT Codex** انجام شده است.

---

## 📘 English README

نسخه کامل انگلیسی مستندات در فایل زیر قرار دارد:

- [`README_EN.md`](./README_EN.md)

</div>
