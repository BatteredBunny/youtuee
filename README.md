
# Youtuee
Funny proxy service that impersonates youtube while at same time redirecting everything to rick roll

```bash
nix run git+https://forge.catnip.ee/batteredbunny/youtuee
```

# Hosting
By default it will try to use ``yt-dlp``, theres an option to call the official api if you include the ``YT_API`` env var, [more info on youtube api keys](https://developers.google.com/youtube/v3/getting-started)

## Nixos

```nix
# flake.nix
youtuee = {
    url = "git+https://forge.catnip.ee/batteredbunny/youtuee";
    inputs.nixpkgs.follows = "nixpkgs";
};
```

```nix
# configuration.nix
imports = [
    inputs.youtuee.nixosModules.default
];

services = {
    youtuee = {
        enable = true;
        user = "youtuee";
        group = "youtuee";

        settings = {
            secretsFile = "/etc/secrets/youtuee.env";
            behindReverseProxy = true;
            port = 8181;
        };
    };

    caddy.virtualHosts."youtu.ee".extraConfig = ''
        reverse_proxy :${toString config.services.youtuee.settings.port}
    '';
};
```