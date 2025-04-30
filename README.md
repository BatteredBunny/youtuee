
# Youtuee
Funny proxy service that impersonates youtube while at same time redirecting everything to rick roll

```bash
nix run git+https://forge.catnip.ee/batteredbunny/youtuee
```

## Hosting on nixos

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

    settings = {
        behindReverseProxy = true;
        port = 8181;
    };
};

caddy.virtualHosts."youtu.ee".extraConfig = ''
    reverse_proxy :${toString config.services.youtuee.settings.port}
'';
};
```