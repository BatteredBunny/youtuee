{ pkgs
, config ? pkgs.config
, lib ? pkgs.lib
, ...
}:
let
  cfg = config.services.youtuee;
in
{
  options.services.youtuee = {
    enable = lib.mkEnableOption "youtuee";

    package = lib.mkOption {
      description = "package to use";
      default = pkgs.callPackage ./build.nix { };
    };

    settings = {
      port = lib.mkOption {
        type = lib.types.int;
        description = "Port to run http api on";
        default = 8080;
      };

      behindReverseProxy = lib.mkEnableOption "Enable if setting up the service behind a reverse proxy" // { default = false; };
    };
  };

  config = lib.mkIf cfg.enable {
    systemd.services.youtuee = {
      enable = true;
      serviceConfig = {
        DynamicUser = true;
        ProtectSystem = "full";
        ProtectHome = "yes";
        DeviceAllow = [ "" ];
        LockPersonality = true;
        MemoryDenyWriteExecute = true;
        PrivateDevices = true;
        ProtectClock = true;
        ProtectControlGroups = true;
        ProtectHostname = true;
        ProtectKernelLogs = true;
        ProtectKernelModules = true;
        ProtectKernelTunables = true;
        ProtectProc = "invisible";
        RestrictNamespaces = true;
        RestrictRealtime = true;
        RestrictSUIDSGID = true;
        SystemCallArchitectures = "native";
        PrivateUsers = true;
        ExecStart = "${lib.getExe cfg.package} -port=${toString cfg.settings.port} ${lib.optionalString cfg.settings.behindReverseProxy "-reverse-proxy"}";
        Restart = "always";
      };
      wantedBy = [ "default.target" ];
    };
  };
}
