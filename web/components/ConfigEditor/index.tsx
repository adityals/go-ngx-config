import Editor from '@monaco-editor/react';

interface Props {
  value: string;
  onChange: (v: string | undefined) => void;
}

const ConfigEditor = (props: Props) => {
  return (
    <div className="h-screen relative block">
      <Editor
        value={props.value}
        onChange={props.onChange}
        className="h-full w-full"
        language="nginx"
        theme="vs-dark"
        defaultValue="// Type your config"
      />
    </div>
  );
};

export default ConfigEditor;
