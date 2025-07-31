{ buildGoModule, pkgs, lib }:
buildGoModule {
  src = ./.;

  name = "youtuee";
  vendorHash = "sha256-V9619+4Fsc3pqFEooB8KioXOYWv5kzGqKTqNvzrbwpA=";

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
