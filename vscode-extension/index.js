const vscode = require('vscode');
const { LanguageClient, LanguageClientOptions, ExitNotification } = require('vscode-languageclient');

function activate(context) {
    let serverOptions = {
        command: '/home/sumeet/goylang/lsp.sh',
        args: []
    };  
    let clientOptions = {
        documentSelector: [{ scheme: 'file', language: 'goylang' }],
        synchronize: {
            configurationSection: 'goylang',
            fileEvents: vscode.workspace.createFileSystemWatcher('**/.goy')
        }
    };
    let client = new LanguageClient('goylang', 'glsp', serverOptions, clientOptions);
    client.start();
}

function deactivate() {
}


module.exports = {
    activate,
    deactivate
}
