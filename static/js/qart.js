function render(operation) {
    fetch('/v1/render', {
        method: 'POST',
        body: JSON.stringify(operation)
    }).then(function (response) {
        return response.json();
    }).then(function (response) {
        if (!response.success) {
            alert(response.message);
            return;
        }
        const element = document.getElementById('op-qr-code');
        if (element) {
            element.src = response.data.image;
        }
    })
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
            element.addEventListener('input', handler);
        } else {
            element.addEventListener('change', handler);
        }
    });
    document.getElementById('op-refresh').addEventListener('click', function (event) {
        render(operation);
    })
})();

