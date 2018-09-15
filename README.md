# Avatargen

Are you running some kind of webservice where users can upload a profile picture but many are missing? Are you tired of looking at the same default avatar? Good, you are in the right place!

Based on a random walk, `avatargen` will create interesting placeholder avatars for you!

## Installation

`go get github.com/doctordesh/avatargen`

## Usage

Run the `avatargen` from the command line.

```
Usage: avatargen [ <flags-and-options> ] <output-directory>

  -alpha float
    	Alpha for the right side of the icon. (default 0.9)
  -color-background string
    	Background color, format <rrggbb> in hex. (default "000000")
  -color-foreground string
    	Foreground color, format <rrggbb> in hex. (default "FFFFFF")
  -count int
    	Number of images to generate. (default 100)
  -scale int
    	The scale of the output image. An icon with 'size' 32 and 'scale' 1 will measure 32x32 pixels. An icon with 'size' 32 and 'scale' 2 will measure 64x64 pixels. (default 5)
  -size int
    	The size of the icon (not in pixels, see 'scale'). (default 32)
  -steps int
    	Number of steps to take when drawing the image. (default 100)
```

## Examples

### White on Black in size 16

`$ avatargen -size 16 -scale 10 images`

![alt text](./doc/example-16-1.png)
![alt text](./doc/example-16-2.png)
![alt text](./doc/example-16-3.png)
![alt text](./doc/example-16-4.png)
![alt text](./doc/example-16-5.png)

### Black on White in size 32 without transparency

`$ avatargen -size 20 -color-background ffffff -color-foreground 000000 -alpha 1 -steps 200 images`

![alt text](./doc/example-20-1.png)
![alt text](./doc/example-20-2.png)
![alt text](./doc/example-20-3.png)
![alt text](./doc/example-20-4.png)
![alt text](./doc/example-20-5.png)

### Specified background and foreground color in size 8

`$ avatargen -size 8 -scale 20 -steps 30 -color-foreground CA5D00 -color-background 124853 images`

![alt text](./doc/example-8-1.png)
![alt text](./doc/example-8-2.png)
![alt text](./doc/example-8-3.png)
![alt text](./doc/example-8-4.png)
![alt text](./doc/example-8-5.png)


