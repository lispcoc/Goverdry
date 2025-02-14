function BoolArrayToBinStr (a, si = 0, ei = 0) {
  let start = si / 16
  let end = ei ? ei / 16 : a.length / 16
  let word = []
  let str = ''
  for (let i = start; i < end; i++) {
    word[i - start] = 0
    for (let j = 0; j < 16; j++) {
      if (a[i * 16 + j]) {
        word[i - start] |= 1 << j
      }
    }
  }
  for (let i = start; i < end; i++) {
    str += word[i - start].toString(2).padStart(16, '0')
  }
  return str
}

function toRGB (s) {
  if (s == undefined || s == null) {
    return { r: 255, g: 255, b: 255 }
  }
  if (s == 'red') {
    s = '#ff0000'
  }
  if (s == 'green') {
    s = '#00ff00'
  }
  if (s == 'blue') {
    s = '#0000ff'
  }
  if (s == 'pink') {
    s = '#ffc0cb'
  }
  if (s == 'yellow') {
    s = '#FFFF00'
  }
  if (s == 'black') {
    s = '#000000'
  }
  if (s == 'white') {
    s = '#ffffff'
  }
  if (s[0] != '#') {
    s = '#000000'
    console.log('Invalid colercode.')
  }
  const [r, g, b] = s.replace('#', '').match(/.{2}/g)
  const red = parseInt(r, 16)
  const green = parseInt(g, 16)
  const blue = parseInt(b, 16)
  return { r: red, g: green, b: blue }
}
