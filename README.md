# Fizz-Image [![CircleCI](https://circleci.com/gh/jschaefer-io/fizz-image.svg?style=svg)](https://circleci.com/gh/jschaefer-io/fizz-image)
Fizz-Image is a custom placeholder image generator.
The images are generated on the fly and can therefore be very easily personalized via the request URL.
It is possible to set the dimensions, background color, text color and label text.

Get started here **`https://fizz-image.herokuapp.com/`**

## Requests
### Dimensions
The minimal required configuration are the image dimensions. \
**Schema: `/{width}x{height}`**

`/100x100` \
![/100x100](https://fizz-image.herokuapp.com/100x100)

`/300x80` \
![/300x80](https://fizz-image.herokuapp.com/300x80)

#### Constraints
- The maximum number of pixels allowed is `4000x4000`

### Background color
The next possible configuration is setting the background color using a hexadecimal color \
**Schema: `/{width}x{height}/{background-color}`**

`/100x100/ff0` \
![/100x100/ff0](https://fizz-image.herokuapp.com/100x100/ff0)

`/300x80/252525` \
![/300x80/252525](https://fizz-image.herokuapp.com/300x80/252525)

#### Constraints
- Only hex-codes with 3 or 6 characters are allowed

### Label color
The next possible configuration is setting the label text color using a hexadecimal color \
**Schema: `/{width}x{height}/{background-color}/{label-color}`**

`/100x100/ff0/000` \
![/100x100/ff0/000](https://fizz-image.herokuapp.com/100x100/ff0/000)

`/300x80/252525/f00` \
![/300x80/252525/f00](https://fizz-image.herokuapp.com/300x80/252525/f00)

#### Constraints
- Only hex-codes with 3 or 6 characters are allowed

### Label Text
Changing the text-content is always possible \
**Schema: `/{schema}?label=Hello+World`**

`/100x100?label=Hello+World` \
![/100x100?label=Hello+World](https://fizz-image.herokuapp.com/100x100?label=Hello+World)

`/300x80/252525?label=Hello+GitHub` \
![/300x80/252525?label=Hello+GitHub](https://fizz-image.herokuapp.com/300x80/252525?label=Hello+GitHub)

#### Constraints
- Chars are limited to the GOs [basicfont](https://godoc.org/golang.org/x/image/font/basicfont)