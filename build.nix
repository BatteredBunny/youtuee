{ buildGoModule, pkgs, lib }:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-c7LLqcxnQXz28nPu/P5aGClC6O65MVzAhKIO1rnFiU8=";

  ldflags = [
    "-s"
    "-w"
  ];

  nativeBuildInputs = with pkgs; [
    makeWrapper
  ];

  postInstall = ''
    wrapProgram $out/bin/youtuee --prefix PATH : ${lib.makeBinPath [ pkgs.yt-dlp ]}
  '';

  meta = {
    homepage = "https://github.com/BatteredBunny/youtuee";
    mainProgram = "youtuee";
  };
}
