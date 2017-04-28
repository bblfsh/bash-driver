-include .sdk/Makefile

$(if $(filter true,$(sdkloaded)),,$(error You must install bblfsh-sdk))

test-native-internal:
	cd native; \
		make; \
		gradle test

build-native-internal:
	cd native; \
		make; \
		gradle installDist;
	echo '#!/usr/bin/env bash\nDIR="$$( cd "$$( dirname "$${BASH_SOURCE[0]}" )" && pwd )"\ncd $${DIR}/../native\nbuild/install/bashdriver/bin/bashdriver' > build/native
	chmod u+x build/native
