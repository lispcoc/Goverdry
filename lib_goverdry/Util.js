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
