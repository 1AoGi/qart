<!DOCTYPE html>

<html dir="ltr">
<head>
    <title>QArt Coder</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/css/bootstrap.min.css" />
    <style>
        .qr-container {
            margin-top: 3em;
            text-align: center;
        }

        .qr-container #op-qr-code, .qr-container .figure {
            width: 100%;
            max-width: 300px;
        }

        .parameter-container {
            margin-top: 1em;
        }
    </style>
 </head>

<body>
    <div class="container">
        <div class="row">
            <div class="qr-container col-sm-6 col-md-4 col-lg-3">
                <figure class="figure">
                    <img id="op-qr-code" src="/image/placeholder/800x800" class="figure-img img-fluid rounded" alt="QR Code">
                </figure>
                <div class="row g-3">
                    <div class="col-6">
                        <button id="op-refresh" type="button" class="btn btn-primary">Refresh</button>
                    </div>
                    <div class="col-6">
                        <button type="button" class="btn btn-primary">Share</button>
                    </div>
                </div>
            </div>
            <div class="parameter-container col-sm-6 col-md-8 col-lg-9">
                <h1>QArt Coder</h1>
                <form class="row g-3" id="parameter-form">
                    <div class="input-group mb-3 col-md-16">
                        <button id="op-upload" type="button" class="btn btn-primary file">
                            Upload
                        </button>
                        <input type="file" id="op-upload-input" multiple accept="image/*" style="display:none">
                        <input class="form-control" type="text" value="https://example.com" id="op-url">
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-rand-control">
                        <label class="form-check-label" for="op-rand-control">Random Pixels</label>
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-only-data-bits">
                        <label class="form-check-label" for="op-only-data-bits">Data Pixels Only</label>
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-dither">
                        <label class="form-check-label" for="op-dither">Dither</label>
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-save-control">
                        <label class="form-check-label" for="op-save-control">Show Controllable Pixels</label>
                    </div>
                    <div class="col-6">
                        <label for="op-dx" class="form-label">X</label>
                        <input type="range" data-reverse class="form-range col-12" min="-50" max="50" step="1" value="-4" id="op-dx">
                        <label for="op-dy" class="form-label">Y</label>
                        <input type="range" data-reverse class="form-range col-12" min="-50" max="50" step="1" value="-4" id="op-dy">
                    </div>
                    <div class="col-6">
                        <label for="op-version" class="form-label">Qr Size</label>
                        <input type="range" class="form-range col-12" min="1" max="9" step="1" value="6" id="op-version">
                        <label for="op-size" class="form-label">Image Size</label>
                        <input type="range" class="form-range col-12" min="-20" max="20" step="1" value="0" id="op-size">
                    </div>
                </form>
            </div>
        </div>
    </div>

    <script src="/static/js/bootstrap.bundle.min.js"></script>
    <script src="/static/js/qart.js"></script>
    <!--script src="/static/js/reload.min.js"></script-->
</body>
</html>
