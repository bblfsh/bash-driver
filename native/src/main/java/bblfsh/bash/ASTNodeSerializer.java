package bblfsh.bash;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import com.intellij.lang.ASTNode;
import com.intellij.psi.tree.TokenSet;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Set;
import java.util.HashSet;
import java.util.stream.Stream;
import java.util.stream.Collectors;

/**
 * Custom Jackson serializer for com.intellij.lang.ASTNode;
 */
public class ASTNodeSerializer extends StdSerializer<ASTNode> {

    private final Set<String> SKIPTOKENS = Stream.of("FILE", "include-command", "for shellcommand", "logical block",
        "[Bash] case pattern", "case pattern", "[Bash] case pattern list", "var-def-element", "function-def-element",
        "group element", "if shellcommand", "conditional shellcommand", "until loop", "simple-command", "[Bash] combined word",
        "while loop", "[Bash] generic bash command", "[Bash] named symbol", "var-use-element").collect(Collectors.toCollection(HashSet::new));

    private final Set<String> SKIPNODES = Stream.of("WHITE_SPACE", "[Bash] linefeed", "[Bash] string begin",
        "[Bash] ;", "[Bash] (", "[Bash] {", "[Bash] )", "[Bash] }", "[Bash] fi", "fi", "[Bash] shebang element",
        "[Bash] if", "[Bash] for", "[Bash] in", "[Bash] do", "[Bash] done", "[Bash] esac", "[Bash] [ (left conditional)",
        "[Bash]  ] (right conditional)", "[Bash] while", "[Bash] case", "[Bash] string end",
        "[Bash] until").collect(Collectors.toCollection(HashSet::new));

    private final Set<String> ADOPTNEXT = Stream.of("[Bash] then", "[Bash] else", "[Bash] elif")
        .collect(Collectors.toCollection(HashSet::new));

    public ASTNodeSerializer() {
        this(null);
    }

    public ASTNodeSerializer(Class<ASTNode> t) {
        super(t);
    }

    @Override
    public void serialize(ASTNode root, JsonGenerator jG, SerializerProvider provider) throws IOException {
        serializeWithChild(root, jG, provider, null);
    }

    public void serializeWithChild(ASTNode root, JsonGenerator jG, SerializerProvider provider, ASTNode addChild) throws IOException {
        final String type = root.getElementType().toString();
        final String text = root.getText();

        jG.writeStartObject();

        jG.writeStringField("@type", type);
        // Some higher level nodes would write everything on the token without this
        if (!SKIPTOKENS.contains(type)) {
            jG.writeStringField("@token", text);
        }

        final int start = root.getStartOffset();
        jG.writeNumberField("startOffset", start);

        final int length = root.getTextLength();
        jG.writeNumberField("endOffset", start + length);


        jG.writeFieldName("children");
        jG.writeStartArray();
        serializeChildren(root.getChildren(null), jG, provider);

        if (addChild != null) {
            serialize(addChild, jG, provider);
        }

        jG.writeEndArray();
        jG.writeEndObject();
    }

    private void serializeChildren(ASTNode[] children, JsonGenerator jG, SerializerProvider provider) throws IOException {
        ArrayList<ASTNode> filteredChildren = new ArrayList<ASTNode>();

        // Remove some useless nodes to unpolute the AST
        for (ASTNode child: children) {
            final String childType = child.getElementType().toString();
            // Skip some redundant nodes (always followed by another more significative one)
            if (SKIPNODES.contains(childType) ||
                (childType.equals("[Bash] generic bash command") && child.getText().equals("source"))) {
                continue;
            }
            filteredChildren.add(child);
        }

        int i = 0;
        while (i < filteredChildren.size())
        {
            ASTNode child = filteredChildren.get(i);
            ASTNode blockChild = null;

            // Bash's AST gives some block nodes not as children of the semantically significative one
            // but as "next one", this fixes it
            if (ADOPTNEXT.contains(child.getElementType().toString()))
            {
                // Reparent the i+1 children to this node
                blockChild = filteredChildren.get(i+1);
                ++i;
            }

            serializeWithChild(child, jG, provider, blockChild);
            ++i;
        }
    }

}
