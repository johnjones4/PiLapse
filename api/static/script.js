(() => {
  const timescale = [
    '5s',
    '10s',
    '30s',
    '1m',
    '1m30s',
    '2m',
    '5m',
    '10m',
    '30m',
    '1h',
    '2h',
    '3h',
    '4h',
    '5h',
    '6h',
    '12h',
    '24h',
    '48h',
    '72h'
  ]

  const sliders = ['interval', 'limit']
  sliders.forEach(id => {
    const slider = document.getElementById(id)
    if (slider) {
      const view = document.getElementById(`${id}-view`)
      slider.setAttribute('min', 0)
      slider.setAttribute('max', timescale.length - 1)
      const i = timescale.findIndex(t => t == view.value)
      if (i >= 0) {
        slider.value = i
      }
      const update = (event) => {
        view.value = timescale[event.target.value]
      }
      slider.addEventListener('input', update)
      slider.addEventListener('change', update)
    }
  })

  if (document.getElementById('table')) {
    setTimeout(() => {
      window.location.assign(window.location.href)
    }, 5000)
  }
})()
