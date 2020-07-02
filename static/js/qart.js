function errorHandler(response) {
    if (response && !response.success) {
        alert(response.message);
    }
    return response
}

function render(operation) {
    return fetch('/v1/render', {
        method: 'POST',
        body: JSON.stringify(operation)
    }).then(function (response) {
        return response.json();
    }).then(errorHandler).then(function (response) {
        const element = document.getElementById('op-qr-code');
        if (element) {
            element.src = response.data.image;
        }
    });
}

function upload(file) {
    let data = new FormData();
    data.append('image', file);

    return fetch('/v1/render/upload', {
        method: 'POST',
        body: data
    }).then(function (response) {
        return response.json();
    }).then(errorHandler);
}

function share(operation) {
    return fetch('/v1/share', {
        method: 'POST',
        body: JSON.stringify({
            image: operation.image
        })
    }).then(function (response) {
        return response.json();
    }).then(errorHandler);
}

function updateOperation(element, obj) {
    if (!obj || !element.id || !element.id.startsWith('op-')) {
        return;
    }
    let id = element.id.replace('op-', '').replace(/-/g, '').toLocaleLowerCase();
    if (!(id in obj)) {
        return;
    }
    switch (element.type) {
        case 'text': {
            obj[id] = element.value;
            break;
        }
        case 'checkbox': {
            obj[id] = element.checked;
            break;
        }
        case 'range': {
            let value = parseInt(element.value, 10);
            obj[id] = 'reverse' in element.dataset ? -value : value;
        }
    }
    render(obj);
}

(function () {
    const operation = {
        image: "default",
        dx: 4,
        dy: 4,
        size: 0,
        url: "https://example.com",
        version: 6,
        mask: 2,
        randcontrol: false,
        dither: false,
        onlydatabits: false,
        savecontrol: false,
        seed: "",
        scale: 4,
        rotation: 0
    };
    document.querySelectorAll('input').forEach(function(element) {
        let handler = function(event) {
            updateOperation(event.target, operation);
        };
        if (element.type === 'text') {
            element.addEventListener('input', handler, false);
        } else {
            element.addEventListener('change', handler);
        }
    });
    document.getElementById('op-refresh').addEventListener('click', function (event) {
        render(operation);
    }, false);
    let uploadInput = document.getElementById('op-upload-input');
    uploadInput.addEventListener('change', function(event) {
        let files = event.target.files;
        if (files && files.length > 0) {
            upload(files[0]).then(function (response) {
                operation.image = response.data.id;
                render(operation);
            });
        }
    }, false);
    document.getElementById('op-qr-code').addEventListener('click', function (event) {
        operation.rotation = (operation.rotation + 1) % 4;
        render(operation);
    }, false);
    document.getElementById('op-upload').addEventListener('click', function (event) {
        uploadInput.click();
    }, false);
    document.getElementById('op-share').addEventListener('click', function (event) {
        share(operation).then(function (response) {
            window.open(`/share/${response.data.id}`);
        });
    });
})();
