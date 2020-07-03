<!DOCTYPE html>

<html dir="ltr" lang="{{.CurLang.Lang}}">
<head>
    <title>QArt Coder</title>
    <meta http-equiv="content-type" content="text/html; charset=utf-8" />
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

        #op-qr-code {
            cursor: pointer;
        }
    </style>
 </head>

<body>
    <div class="container">
        <div class="row">
            <div class="qr-container col-sm-6 col-md-4 col-lg-3">
                <figure class="figure">
                    <img id="op-qr-code"
                         class="figure-img img-fluid rounded"
                         src="/image/placeholder/800x800" alt="QR Code"
                         data-toggle="tooltip" data-placement="top" title='{{i18n .Lang "index.rotate"}}'
                    />
                </figure>
                <div class="row g-3">
                    <div class="col-6">
                        <button id="op-refresh" type="button" class="btn btn-primary">{{i18n .Lang "index.refresh"}}</button>
                    </div>
                    <div class="col-6">
                        <button id="op-share" type="button" class="btn btn-primary">{{i18n .Lang "index.share"}}</button>
                    </div>
                </div>
            </div>
            <div class="parameter-container col-sm-6 col-md-8 col-lg-9">
                <div class="row">
                    <h1 class="col">QArt Coder</h1>
                    <div class="dropdown col">
                        <button class="btn btn-secondary dropdown-toggle" type="button" id="dropdownMenuButton" data-toggle="dropdown" aria-expanded="false">
                            {{.CurLang.Name}}
                        </button>
                        <ul class="dropdown-menu" aria-labelledby="dropdownMenuButton">
                            {{range $i, $v := .RestLangs}}
                                <li><a class="dropdown-item" href="/?lang={{$v.Lang}}">{{$v.Name}}</a></li>
                            {{end}}
                        </ul>
                    </div>
                </div>
                <form class="row g-3" id="parameter-form">
                    <div class="input-group mb-3 col-md-16">
                        <button id="op-upload" type="button" class="btn btn-primary file">
                            {{i18n .Lang "index.upload"}}
                        </button>
                        <input type="file" id="op-upload-input" accept="image/*" style="display:none" />
                        <input class="form-control" type="text" value="https://example.com" id="op-url" />
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-rand-control" />
                        <label class="form-check-label" for="op-rand-control">{{i18n .Lang "index.rand_control"}}</label>
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-only-data-bits" />
                        <label class="form-check-label" for="op-only-data-bits">{{i18n .Lang "index.only_data_bits"}}</label>
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-dither" />
                        <label class="form-check-label" for="op-dither">{{i18n .Lang "index.dither"}}</label>
                    </div>
                    <div class="form-check form-switch col-md-6">
                        <input class="form-check-input" type="checkbox" value="" id="op-save-control" />
                        <label class="form-check-label" for="op-save-control">{{i18n .Lang "index.save_control"}}</label>
                    </div>
                    <div class="col-6">
                        <label for="op-dx" class="form-label">X</label>
                        <input type="range" data-reverse="" class="form-range col-12" min="-50" max="50" step="1" value="-4" id="op-dx" />
                        <label for="op-dy" class="form-label">Y</label>
                        <input type="range" data-reverse="" class="form-range col-12" min="-50" max="50" step="1" value="-4" id="op-dy" />
                    </div>
                    <div class="col-6">
                        <label for="op-version" class="form-label">{{i18n .Lang "index.qr_version"}}</label>
                        <input type="range" class="form-range col-12" min="1" max="9" step="1" value="6" id="op-version" />
                        <label for="op-size" class="form-label">{{i18n .Lang "index.image_size"}}</label>
                        <input type="range" class="form-range col-12" min="-20" max="20" step="1" value="0" id="op-size" />
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
