import { useLexicalComposerContext } from '@lexical/react/LexicalComposerContext';
import { LexicalEditor, $insertNodes, TextNode, $getRoot, $getSelection } from 'lexical';
import { DateTimeNode } from '../../nodes/DateTimeNode';
import './index.css';

interface DateTimeProps {
  date?: Date;
  onCalendarIconClick?: () => void;
}

function DateTimeComponent({
  date,
  onCalendarIconClick,
}: DateTimeProps): JSX.Element {
  return (
    <div className="date-time-plugin">
      <button
        className="ui button"
        type="button"
        aria-label="insert date and time"
        title="insert date and time"
        onClick={onCalendarIconClick}
      >
        <i className="icon calendar" />
      </button>
      {date?.toLocaleString()}
    </div>
  );
}

function useDateTimeComponent(editor: LexicalEditor, date: Date): JSX.Element {
  const onCalendarIconClick = () => {
    editor.update(() => {
      const node = new DateTimeNode(date.toLocaleString())
      const nodes = [node, new TextNode(" ")]

      const selection = $getSelection();
      if (selection?.getTextContent()) {
        console.log("Text Selected", selection.getTextContent())
        selection.insertNodes(nodes)
      } else {
      $insertNodes(nodes)
      }
    })
  }

  return <DateTimeComponent onCalendarIconClick={onCalendarIconClick} date={date} />
}


export function  DateTimePlugin({
  date = new Date(),
}: DateTimeProps): JSX.Element {
  const [editor] = useLexicalComposerContext();
  return useDateTimeComponent(editor, date);
}
