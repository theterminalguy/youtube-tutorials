import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { PlainTextPlugin } from "@lexical/react/LexicalPlainTextPlugin";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";

import "./Editor.css";
import { AutoFocusPlugin } from "./plugins/AutoFocusPlugin";
import TreeViewPlugin from "./plugins/TreeViewPlugin";
import { DateTimePlugin } from './plugins/DateTimePlugin/index';
import { DateTimeNode } from "./nodes/DateTimeNode";

const theme = {
  // Theme styling goes here
  // ...
};


// Catch any errors that occur during Lexical updates and log them
// or throw them as needed. If you don't throw them, Lexical will
// try to recover gracefully without losing user data.
function onError(error: Error) {
  console.error(error);
}

export function Editor() {
  const initialConfig = {
    namespace: "MyEditor",
    theme,
    onError,
    nodes: [
      DateTimeNode
    ]
  };

  return (
    <LexicalComposer initialConfig={initialConfig}>
        <DateTimePlugin date={new Date()}/>
      <div className="editor-container">
        <div className="editor-inner">
          <PlainTextPlugin
            contentEditable={<ContentEditable className="editor-input"/>}
            placeholder={<div className="editor-placeholder">Enter some text...</div>}
            ErrorBoundary={LexicalErrorBoundary}
          />
          <HistoryPlugin />
          <AutoFocusPlugin />
          <TreeViewPlugin />
        </div>
      </div>
    </LexicalComposer>
  );
}
