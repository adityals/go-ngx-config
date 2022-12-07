import Editor from '@monaco-editor/react';

interface Props {
  value: string;
  languange: string;
}

const ParserResult = (props: Props) => {
  return (
    <Editor
      value={props.value}
      className="h-full w-full"
      language={props.languange}
      theme="vs-dark"
      defaultValue="// Type your config"
    />
  );
};

export default ParserResult;
