# Gixel Maimai

Maimai clone using Gixel-Engine.

## How to build for Android

```sh
git clone https://github.com/odedro987/gixel-maimai.git
cd gixel-maimai
go run github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile bind -target android -androidapi 19 -javapkg gixel -o ./mobile/android/maimai/maimai.aar ./maimai
```

and run the Android Studio project in ./mobile/android.
