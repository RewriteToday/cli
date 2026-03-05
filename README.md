<div align="center">

# Rewrite CLI

The fastest way to build and operate SMS flows with Rewrite.
From login to live logs, stay in flow and ship faster from the command line.

[Website](https://rewritetoday.com) • [Dashboard](https://dash.rewritetoday.com) • [Docs](https://docs.rewritetoday.com) • [CLI Docs](https://docs.rewritetoday.com/en/cli)

<img src="https://cdn.rewritetoday.com/assets/banners/cli.png" width="100%" alt="Rewrite CLI Banner"/>

## Install

</div>

```sh
curl -fsSL https://rewritetoday.com/install | bash
```

```sh
rewrite -v
```

<div align="center">

## Syntax

Installing the CLI provides access to the `rewrite` command

</div>

```bash
rewrite [command] [...flags] [...args]
```

<div align="center">

## Starting with the CLI

Before you enjoy what we have to offer, it is highly recommended to connect your account to the CLI

</div>

```bash
rewrite login
```

<div align="center">

After authentication (*OAuth Device Flow*), your account is automatically connected and you can use everything you want.

## Storing

Your account token is securely stored in your operating system's native keyring, so you do not need to worry about it.

If the token cannot be saved, you must install the keyring on your operating system (Linux only)

</div>

```bash
sudo apt install gnome-keyring
```

<div align="center">

Or any other keyring based on your distro.

Use `rewrite <command> -h` for command help.

For complete command reference and examples:
[https://docs.rewritetoday.com/en/cli](https://docs.rewritetoday.com/en/cli)

---

Made with 🤍 by the Rewrite team. <br/>
SMS the way it should be.

</div>
