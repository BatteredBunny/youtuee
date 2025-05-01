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

    user = lib.mkOption {
      type = lib.types.str;
      default = "youtuee";
      description = "User account under which youtuee runs.";
    };

    group = lib.mkOption {
      type = lib.types.str;
      default = "youtuee";
      description = "Group under which youtuee runs.";
    };

    settings = {
      port = lib.mkOption {
        type = lib.types.int;
        description = "Port to run http api on";
        default = 8080;
      };

      behindReverseProxy = lib.mkEnableOption "Enable if setting up the service behind a reverse proxy" // { default = false; };

      secretsFile = lib.mkOption {
        type = lib.types.nullOr lib.types.path;
        default = null;
        description = ''
          Allows specifying YT_API
        '';
      };
    };
  };

  config = lib.mkIf cfg.enable {
    users.users = lib.mkIf (cfg.user == "youtuee") {
      youtuee = {
        group = cfg.group;
        createHome = false;
        isSystemUser = true;
      };
    };

    users.groups.${cfg.group} = { };

    systemd.services.youtuee = {
      enable = true;
      serviceConfig = {
        User = cfg.user;
        Group = cfg.group;
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
        EnvironmentFile = cfg.settings.secretsFile;
        Restart = "always";
      };

      environment.GIN_MODE = "release";
      wantedBy = [ "default.target" ];
    };
  };
}
