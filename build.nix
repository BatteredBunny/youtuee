{ buildGoModule }:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-pIwL77oLQfxipbTcTJNBwZ14dZvpDblOrNBZjcjYIZA=";

  ldflags = [
    "-s"
    "-w"
  ];

  meta = {
    homepage = "https://github.com/BatteredBunny/youtuee";
    mainProgram = "youtuee";
  };
}
