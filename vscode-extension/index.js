const vscode = require('vscode');
const { LanguageClient, LanguageClientOptions, ExitNotification, RevealOutputChannelOn } = require('vscode-languageclient');

function activate(context) {
    let serverOptions = {
        command: '/home/sumeet/goylang/lsp-debug.sh',
        args: []
    };
    let clientOptions = {
        documentSelector: [{ scheme: 'file', language: 'goylang' }],
        revealOutputChannelOn: RevealOutputChannelOn.Info,
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
