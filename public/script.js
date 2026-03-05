async function loadData(){

const res = await fetch("/api/images")
const data = await res.json()

document.getElementById("largeSize").innerText = data.largeImage.size
document.getElementById("largeBuild").innerText = data.largeImage.buildTime
document.getElementById("largeStart").innerText = data.largeImage.startTime

document.getElementById("smallSize").innerText = data.optimizedImage.size
document.getElementById("smallBuild").innerText = data.optimizedImage.buildTime
document.getElementById("smallStart").innerText = data.optimizedImage.startTime

}