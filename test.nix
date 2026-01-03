{ self, testers }:
testers.nixosTest {
  name = "youtuee";

  nodes.machine =
    { pkgs, ... }:
    {
      imports = [
        self.nixosModules.default
      ];

      services.youtuee = {
        enable = true;
        settings.port = 8888;
      };

      environment.systemPackages = with pkgs; [
        curl
      ];
    };

  testScript =
    { nodes, ... }:
    let
      serviceUrl = "http://localhost:${toString nodes.machine.services.youtuee.settings.port}";
    in
    ''
      start_all()
      machine.wait_for_unit("youtuee.service")
      machine.wait_for_open_port(${toString nodes.machine.services.youtuee.settings.port})
      machine.succeed("curl ${serviceUrl}")
      machine.succeed("curl ${serviceUrl}/j_fkAFRPaHQ | grep -o 'FX Artists React to Bad &amp; Great CGi 115'")
    '';
}
