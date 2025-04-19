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

  testScript = { nodes, ... }: ''
    start_all()
    machine.wait_for_unit("youtuee.service")
    machine.wait_for_open_port(${toString nodes.machine.services.youtuee.settings.port})
    machine.succeed("curl 0.0.0.0:${toString nodes.machine.services.youtuee.settings.port} | grep -o \"https://github.com/BatteredBunny/youtuee\"")
  '';
}
