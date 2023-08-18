{
  description = "A simple Go package";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs = {
      url = "github:Nixos/nixpkgs/nixpkgs-unstable";
      flake = true;
    };
    flake-compat = {
      url = "github:edolstra/flake-compat";
      flake = false;
    };
  };

  outputs = { self, nixpkgs, flake-compat }:
    let

      # to work with older version of flakes
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";

      # Generate a user-friendly version number.
      version = builtins.substring 0 8 lastModifiedDate;

      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {

      # Provide some binary packages for selected system types.
      packages = forAllSystems
        (system:
          let
            pkgs = nixpkgsFor.${system};
            name = "Seekr";
            appdir = "${name}.AppDir";
          in
          {

            seekr-appimage = pkgs.stdenv.mkDerivation {
              name = "seekr-appimage";

              #src = self.packages.${system}.seekr;
              src = ./.;
              buildInputs = [
                pkgs.appimagekit
                self.packages.${system}.seekr
              ];
              buildPhase = ''
                mkdir -p ${appdir}/usr/bin
                cp web/images/256x256.png ${appdir}/${name}.png
                
                install ${self.packages.${system}.seekr}/bin/seekr ${appdir}/usr/bin/seekr
                cat > ${appdir}/${name}.desktop <<EOF
                [Desktop Entry]
                Type=Application
                Name=Seekr
                Exec=seekr %U
                Icon=Seekr
                StartupNotify=true
                Categories=Network;
                EOF
                cp AppRun ${appdir}/AppRun
                chmod +x ${appdir}/AppRun
                appimagetool -v -n ./Seekr.AppDir ./Seekr.AppImage
                appimagetool -l ./Seekr.AppImage
              '';
              installPhase = ''
                mkdir -p $out/bin
                cp -r ${appdir} $out/${appdir}

                install Seekr.AppImage $out/Seekr.AppImage
              '';

            };

            seekr = pkgs.buildGoModule {
              #${pkgs.git}/bin/git init -q
              postConfigure = ''
                ${pkgs.go}/bin/go generate ./...
                ${pkgs.nodePackages_latest.typescript}/bin/tsc --project web --watch false
              '';
              pname = "seekr";
              inherit version;
              src = ./.;
              CGO_ENABLED = 0;
              tags = [
                "osusergo"
                "netgo"
                "static_build"
              ];
              ldflags = [
                "-s -w"
                "-extldflags=-static"
                #"-X main.version=${version}"
              ];
              vendorSha256 = "sha256-k7PM0XZRMz+gPWKbLy7fJrtlnhUOOGAbOdbSTs8L1y0=";
            };
          });

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = "${self.packages.${system}.seekr}/bin/seekr";
        };

      });

      formatter = forAllSystems (system: nixpkgsFor.${system}.nixpkgs-fmt);

      devShells = forAllSystems (system: {
        default = nixpkgsFor.${system}.mkShell {
          packages = [
            nixpkgsFor.${system}.nodePackages_latest.typescript
            nixpkgsFor.${system}.go
            nixpkgsFor.${system}.gcc
            # jq is useful to debug the database
            nixpkgsFor.${system}.jq
            nixpkgsFor.${system}.gh
            nixpkgsFor.${system}.goreleaser

            nixpkgsFor.${system}.gcc
          ];
        };
        shellHook = ''
          sbuild() {
            go generate ./...
            tsc --project web
            go run main.go
          }
        '';
      });


      defaultPackage = forAllSystems (system: self.packages.${system}.seekr);
    };
}
