const go = new Go();
if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    }
}

WebAssembly
    .instantiateStreaming(fetch('main.wasm'), go.importObject)
    .then(result => {
        go.run(result.instance);
    });

window.onload = () => {
    const uniInput = document.getElementById("uni-input")
    const zgInput = document.getElementById("zg-input")
    const txtFileInput = document.getElementById("txtFileInput")
    //const txtFileResult = document.getElementById("txtFileResult")

    uniInput.onkeyup = (e) => {
        let res = toZg(e.target.value)
        zgInput.value = res.data
    }

    zgInput.onkeyup = (e) => {
        let res = toUni(e.target.value)
        uniInput.value = res.data
    }

    txtFileInput.onchange = (e) => {
        let filename = e.target.files[0].name
        convertTextFile(e.target.files[0]).then(resp => {
            let blob = new Blob([resp.data], {
                type: "text/plain;charset=utf-8"
            })
            const link = document.createElement("a")
            link.href = URL.createObjectURL(blob)
            link.download = `${filename}`
            link.click();
            URL.revokeObjectURL(link.href);
        })
    }
}