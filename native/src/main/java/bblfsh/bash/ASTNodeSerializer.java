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
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Stream;
import java.util.stream.Collectors;

/**
 * Custom Jackson serializer for com.intellij.lang.ASTNode;
 */
public class ASTNodeSerializer extends StdSerializer<ASTNode> {

    private static final Map<String, String> TRANS_TABLE = createTransMap();

    private final Set<String> SKIPTOKENS = Stream.of("FILE", "include-command", "for_shellcommand", "logical_block",
        "case_pattern", "case_pattern_list", "var-def-element", "function-def-element",
        "group_element", "if_shellcommand", "conditional_shellcommand", "until_loop", "simple-command", "combined_word",
        "while_loop", "generic_bash_command", "named_symbol", "var-use-element").collect(Collectors.toCollection(HashSet::new));

    private final Set<String> SKIPNODES = Stream.of("WHITE_SPACE", "linefeed", "string_begin",
        ";", "(", "{", ")", "}", "fi", "fi", "shebang_element",
        "if", "for", "in", "do", "done", "esac", "[_(left_conditional)",
        "]_(right_conditional)", "while", "case", "string_end",
        "until").collect(Collectors.toCollection(HashSet::new));

    private final Set<String> ADOPTNEXT = Stream.of("then", "else", "elif")
        .collect(Collectors.toCollection(HashSet::new));

    private static Map<String, String> createTransMap() {
        Map<String, String> m = new HashMap<String, String>();
        m.put("File_reference", "file_ref");
        m.put("unevaluated_string_(STRING2)", "unevaluated_string2");
        m.put("arith_==", "arith_EQEQ");
        m.put("arith_<", "arith_LT");
        m.put("arith_>", "arith_GT");
        m.put("==", "EQEQ");
        m.put("=", "EQ");
        m.put("arith_!=", "arith_NOTEQ");
        m.put("!=", "NOTEQ");
        m.put("<<", "LT");
        m.put("<", "LT");
        m.put("<=", "LTEQ");
        m.put(">", "GT");
        m.put(">>", "GTGT");
        m.put(">=", "GTEQ");
        m.put("||", "OROR");
        m.put("|", "OR");
        m.put("&&", "ANDAND");
        m.put("&", "AND");
        m.put("cond_op_!", "cond_op_NOT");
        m.put("cond_op_==", "cond_op_EQEQ");
        m.put("conditional_shellcommand", "conditional_shellcommand");
        m.put("[_for_arithmetic", "LB_for_arithmetic");
        m.put("]_for_arithmetic", "RB_for_arithmetic");
        m.put(":", "COLON");
        m.put("[_(left_square)", "LB_left_square");
        m.put("]_(right_square)", "RB_right_square");
        m.put(";;", "SEMICOLONSEMICOLON");
        m.put("[[_(left_bracket)", "LBLB_left_bracket");
        m.put("]]_(right_bracket)", "RBRB_right_bracket");
        m.put("((", "LPLP");
        m.put("))", "RPRP");
        m.put("backquote_`", "backquote");
        m.put("$", "DOLLAR");
        m.put("composed_variable,_like_subshell", "composed_variable");
        m.put("&[0-9]_filedescriptor", "numrange_filedescriptor");
        m.put("Parameter_expansion_operator_'@@'", "param_exp_ATAT");
        m.put("Parameter_expansion_operator_'@'", "param_exp_AT");
        m.put("Parameter_expansion_operator_'##'", "param_exp_NUMNUM");
        m.put("Parameter_expansion_operator_'#'", "param_exp_NUM");
        m.put("Parameter_expansion_operator_'%%'", "param_exp_PERCPERC");
        m.put("Parameter_expansion_operator_'%'", "param_exp_PERC");
        m.put("Parameter_expansion_operator_'::'", "param_exp_COLONCOLON");
        m.put("Parameter_expansion_operator_':'", "param_exp_COLON");
        m.put("Parameter_expansion_operator_'//'", "param_exp_SLASHSLASH");
        m.put("Parameter_expansion_operator_'/'", "param_exp_SLASH");
        m.put("lazy_LET_expression", "lazy_let_expr");
        return m;
    }

    private static String translateType(String type) {
        String t = type.replace("[Bash] ", "").trim().replace(" ", "_");
        return TRANS_TABLE.getOrDefault(t, t);
    }

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
        final String type = translateType(root.getElementType().toString());
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
            final String childType = translateType(child.getElementType().toString());
            // Skip some redundant nodes (always followed by another more significative one)
            if (SKIPNODES.contains(childType) ||
                (childType.equals("generic_bash_command") && child.getText().equals("source"))) {
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
            if (ADOPTNEXT.contains(translateType(child.getElementType().toString())))
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
