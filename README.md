# Fizz-Image
Fizz-Image is a custom placeholder image generator.
The images are generated on the fly and can therefore be very easily personalized via the request URL.
It is possible to set the dimensions, background color, text color and label text.

## Requests
### Dimensions
The minimal required configuration are the image dimensions. \
**Schema: `/{width}x{height}`**

`/100x100` \
![/100x100](https://vast-headland-40276.herokuapp.com/100x100)

`/300x80` \
![/300x80](https://vast-headland-40276.herokuapp.com/300x80)

### Background color
The next possible configuration is setting the background color using a hexadecimal color \
**Schema: `/{width}x{height}/{background-color}`**

`/100x100/ff0` \
![/100x100/ff0](https://vast-headland-40276.herokuapp.com/100x100/ff0)

`/300x80/252525` \
![/300x80/252525](https://vast-headland-40276.herokuapp.com/300x80/252525)

### Label color
The next possible configuration is setting the label text color using a hexadecimal color \
**Schema: `/{width}x{height}/{background-color}/{label-color}`**

`/100x100/ff0/000` \
![/100x100/ff0/000](https://vast-headland-40276.herokuapp.com/100x100/ff0/000)

`/300x80/252525/f00` \
![/300x80/252525/f00](https://vast-headland-40276.herokuapp.com/300x80/252525/f00)

### Label Text
Changing the text-content is always possible \
**Schema: `/{schema}?label=Hello+World`**

`/100x100?label=Hello+World` \
![/100x100?label=Hello+World](https://vast-headland-40276.herokuapp.com/100x100?label=Hello+World)

`/300x80/252525?label=Hello+GitHub` \
![/300x80/252525?label=Hello+GitHub](https://vast-headland-40276.herokuapp.com/300x80/252525?label=Hello+GitHub)