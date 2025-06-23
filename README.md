
# Youtuee
Funny proxy service that impersonates youtube while at same time redirecting everything to rick roll

To use simply, click share on a video then replace be with ee (youtu.be -> youtu.ee)

![CleanShot 2025-06-23 at 14 16 40@2x](https://github.com/user-attachments/assets/ccfdf53f-4111-43b8-a334-63cac82e0d96)

https://github.com/user-attachments/assets/54e0901b-df1e-43a7-a656-43eda2134364

# Hosting
By default it will try to use ``yt-dlp``, theres an option to call the official api if you include the ``YT_API`` env var, [more info on youtube api keys](https://developers.google.com/youtube/v3/getting-started)

## Nixos

```nix
# flake.nix
youtuee = {
    url = "github:batteredbunny/youtuee";
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

# Development

Runnings tests

```
go test ./cmd -v
nix flake check .
```
