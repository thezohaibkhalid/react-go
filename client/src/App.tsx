import { Blockquote } from "@chakra-ui/react";

export default function App() {
  return (
    <div>
      <Blockquote.Root>
        <Blockquote.Content cite="https://" />
        <Blockquote.Caption>
          <cite>Uzumaki Naruto</cite>
        </Blockquote.Caption>
      </Blockquote.Root>
    </div>
  );
}
