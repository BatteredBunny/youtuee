{ nixosTest, nixosModule }:
nixosTest {
  name = "youtuee";
  nodes.machine = { pkgs, ... }: {
    imports = [ nixosModule ];
    services.youtuee.enable = true;

    environment.systemPackages = with pkgs; [
      curl
    ];
  };

  testScript = { nodes, ... }: let
    serviceUrl = "0.0.0.0:${toString nodes.machine.services.youtuee.settings.port}";
  in ''
    start_all()
    machine.wait_for_unit("youtuee.service")
    machine.wait_for_open_port(${toString nodes.machine.services.youtuee.settings.port})
    machine.succeed("curl ${serviceUrl} | grep -o \"https://github.com/BatteredBunny/youtuee\"")
    machine.succeed("curl ${serviceUrl}/j_fkAFRPaHQ | grep -o \"FX Artists React to Bad &amp; Great CGi 115\"")
  '';
}
