import { ThemedText } from "@/components/ThemedText";
import { useCreateEvent, useCurrentBaby } from "@/hooks/api";
import { NewEvent } from "@/types/api";
import { router } from "expo-router";
import { Button, View } from "react-native";

export default function AddEvent() {
  const { data: baby } = useCurrentBaby();
  const mutation = useCreateEvent(baby?.id ?? 0);

  return (
    <View>
      <ThemedText>AddEvent: {JSON.stringify(baby, null, 2)}</ThemedText>
      <Button
        title="Add Event"
        onPress={async () => {
          try {
            await mutation.mutateAsync(makeTestEvent());
            router.back();
          } catch (e) {
            console.error(e);
          }
        }}
      />
      <Button title="Cancel" onPress={() => router.back()} />
    </View>
  );
}

function makeTestEvent(): NewEvent {
  return {
    eventType: "feeding",
    startTime: new Date().toISOString(),
  };
}
