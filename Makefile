-include .sdk/Makefile

$(if $(filter true,$(sdkloaded)),,$(error You must install bblfsh-sdk))

NATIVE_SCRIPT := native/src/main/ash/native.ash
DOWNLOAD_VENDOR = make
BUILD = gradle
JAR := native/build/libs/native-jar-with-dependencies.jar

test-native-internal:
	cd native; \
		$(DOWNLOAD_VENDOR); \
		$(BUILD) test

build-native-internal:
	cd native; \
		$(DOWNLOAD_VENDOR); \
		$(BUILD) shadowJar;
	cp $(JAR) $(BUILD_PATH);
	cp $(NATIVE_SCRIPT) $(BUILD_PATH)/native;
	chmod +x $(BUILD_PATH)/native
