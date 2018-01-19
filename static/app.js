var ap = undefined;
const defaultVolume = 0.5;

function play(e, i) {
    var volume = defaultVolume;
    if (ap != undefined) {
        var volBar = document.getElementsByClassName('aplayer-volume')[0];
        if (volBar != undefined) {
            volume = parseInt(volBar.style.height) / 100;
        }
        ap.destroy();
    }

    ap = new APlayer({
        element: document.getElementById('player'),
        music: albums[i]
    });
    ap.volume(volume);
    ap.play();
}
