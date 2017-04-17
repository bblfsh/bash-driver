package bblfsh.bash;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.SerializerProvider;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import com.intellij.lang.ASTNode;

import java.io.IOException;
import java.util.List;

/**
 * Custom Jackson serializer for com.intellij.lang.ASTNode;
 */
public class ASTNodeSerializer extends StdSerializer<ASTNode> {

    public ASTNodeSerializer() {
        this(null);
    }

    public ASTNodeSerializer(Class<ASTNode> t) {
        super(t);
    }

    @Override
    public void serialize(ASTNode root, JsonGenerator jG, SerializerProvider provider) throws IOException {
        jG.writeStartObject();
        // TODO: actually serialize the parse results into JSON.
        jG.writeEndObject();
    }

}
