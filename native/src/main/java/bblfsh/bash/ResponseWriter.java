package bblfsh.bash;

import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;
import java.io.OutputStream;

public class ResponseWriter {

    private final OutputStream out;
    private final ObjectMapper mapper;

    public ResponseWriter(final OutputStream out, final ObjectMapper mapper) {
        this.out = out;
        this.mapper = mapper;
    }

    public void write(final Response response) throws IOException {
        mapper.writeValue(this.out, response);
        this.out.write('\n');
    }

}
