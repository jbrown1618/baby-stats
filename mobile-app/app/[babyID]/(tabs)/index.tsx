import { ThemedText } from "@/components/ThemedText";
import { View } from "react-native";
import { useCurrentBaby } from "@/hooks/api";

export default function TabsIndex() {
  const { data: baby } = useCurrentBaby();
  return (
    <View>
      <ThemedText>TabsIndex: {JSON.stringify(baby, null, 2)}</ThemedText>
    </View>
  );
}
