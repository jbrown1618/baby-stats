import { ThemedText } from "@/components/ThemedText";
import { useLocalSearchParams } from "expo-router";
import { View } from "react-native";

export default function TabsIndex() {
  const p = useLocalSearchParams();
  return (
    <View>
      <ThemedText>TabsIndex: {JSON.stringify(p)}</ThemedText>
    </View>
  );
}
